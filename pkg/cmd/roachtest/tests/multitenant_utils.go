// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

import (
	"context"
	gosql "database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/cluster"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/option"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/test"
	"github.com/cockroachdb/cockroach/pkg/roachprod"
	"github.com/cockroachdb/cockroach/pkg/roachprod/config"
	"github.com/cockroachdb/cockroach/pkg/security/username"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/stretchr/testify/require"
)

// TODO: move this below cluster interface, maybe into roachprod.

// Map tenantId to instance count so we can uniquify each instance and
// run multiple instances on one node.
var tenantIds = make(map[int]int)

// tenantNode corresponds to a running tenant.
type tenantNode struct {
	tenantID          int
	instanceID        int
	httpPort, sqlPort int
	kvAddrs           []string
	pgURL             string
	// Same as pgURL but with relative ssl parameters.
	relativeSecureURL string

	binary string // the binary last passed to start()
	errCh  chan error
	node   int
}

type createTenantOptions struct {
	// TODO(ssd): This is a hack to work around the currently tangled state of
	// cluster management between roachtest and roachprod. createTenantNode
	// recreates client certs. Only one copy of the client certs are cached
	// locally, so if we want a client to work against multiple tenants in a
	// single test, we need to create the certs with all tenants.
	otherTenantIDs []int

	// Set this to expand the scope of the nodes added to the tenant certs.
	certNodes option.NodeListOption
}
type createTenantOpt func(*createTenantOptions)

func createTenantOtherTenantIDs(ids []int) createTenantOpt {
	return func(c *createTenantOptions) { c.otherTenantIDs = ids }
}

func createTenantCertNodes(nodes option.NodeListOption) createTenantOpt {
	return func(c *createTenantOptions) { c.certNodes = nodes }
}

func createTenantNode(
	ctx context.Context,
	t test.Test,
	c cluster.Cluster,
	kvnodes option.NodeListOption,
	tenantID, node, httpPort, sqlPort int,
	opts ...createTenantOpt,
) *tenantNode {
	var createOptions createTenantOptions
	for _, o := range opts {
		o(&createOptions)
	}
	instID := tenantIds[tenantID] + 1
	tenantIds[tenantID] = instID
	// In secure mode only the internal address works, possibly because its started
	// with --advertise-addr on internal address?
	kvAddrs, err := c.InternalAddr(ctx, t.L(), kvnodes)
	require.NoError(t, err)
	tn := &tenantNode{
		tenantID:   tenantID,
		instanceID: instID,
		httpPort:   httpPort,
		kvAddrs:    kvAddrs,
		node:       node,
		sqlPort:    sqlPort,
	}
	if tn.cockroachBinSupportsTenantScope(ctx, c) {
		err := tn.recreateClientCertsWithTenantScope(ctx, c, createOptions.otherTenantIDs)
		require.NoError(t, err)
	}
	tn.createTenantCert(ctx, t, c, createOptions.certNodes)
	return tn
}

// cockroachBinSupportsTenantScope is a hack to figure out if the version of
// cockroach on the node supports tenant scoped certificates. We can't use a
// version comparison here because we need to compare alpha build versions which
// are compared lexicographically. This is a problem because our alpha versions
// contain an integer count of commits, which does not sort correctly.  Once
// this feature ships in a release, it will be easier to do a version comparison
// on whether this command line flag is supported.
func (tn *tenantNode) cockroachBinSupportsTenantScope(ctx context.Context, c cluster.Cluster) bool {
	err := c.RunE(ctx, c.Node(tn.node), "./cockroach cert create-client --help | grep '\\--tenant-scope'")
	return err == nil
}

func (tn *tenantNode) createTenantCert(
	ctx context.Context, t test.Test, c cluster.Cluster, certNodes option.NodeListOption,
) {

	if len(certNodes) == 0 {
		certNodes = c.Node(tn.node)
	}

	var names []string
	for _, node := range certNodes {
		eips, err := c.ExternalIP(ctx, t.L(), c.Node(node))
		require.NoError(t, err)
		names = append(names, eips...)
		iips, err := c.InternalIP(ctx, t.L(), c.Node(node))
		require.NoError(t, err)
		names = append(names, iips...)
	}
	names = append(names, "localhost", "127.0.0.1")

	cmd := fmt.Sprintf(
		"./cockroach cert create-tenant-client --certs-dir=certs --ca-key=certs/ca.key %d %s",
		tn.tenantID, strings.Join(names, " "))
	c.Run(ctx, c.Node(tn.node), cmd)
}

func (tn *tenantNode) recreateClientCertsWithTenantScope(
	ctx context.Context, c cluster.Cluster, otherIDs []int,
) error {
	tenantArgs := fmt.Sprintf("1,%d", tn.tenantID)
	for _, id := range otherIDs {
		tenantArgs = fmt.Sprintf("%s,%d", tenantArgs, id)
	}

	for _, user := range []username.SQLUsername{username.RootUserName(), username.TestUserName()} {
		cmd := fmt.Sprintf(
			"./cockroach cert create-client %s --certs-dir=certs --ca-key=certs/ca.key --tenant-scope %s --overwrite",
			user.Normalized(), tenantArgs)
		c.Run(ctx, c.Node(tn.node), cmd)
	}
	return c.RefetchCertsFromNode(ctx, tn.node)
}

func (tn *tenantNode) stop(ctx context.Context, t test.Test, c cluster.Cluster) {
	if tn.errCh == nil {
		return
	}
	// Must use pkill because the context cancellation doesn't wait for the
	// process to exit.
	c.Run(ctx, c.Node(tn.node),
		fmt.Sprintf("pkill -o -f '^%s mt start.*tenant-id=%d.*%d'", tn.binary, tn.tenantID, tn.sqlPort))
	t.L().Printf("mt cluster exited: %v", <-tn.errCh)
	tn.errCh = nil
}

func (tn *tenantNode) logDir() string {
	return fmt.Sprintf("logs/mt-%d-%d", tn.tenantID, tn.instanceID)
}

func (tn *tenantNode) storeDir() string {
	return fmt.Sprintf("cockroach-data-mt-%d-%d", tn.tenantID, tn.instanceID)
}

// In secure mode the url we get from roachprod contains ssl parameters with
// local file paths. secureURL returns a url with those changed to
// roachprod/workload friendly local paths, ie "certs".
func (tn *tenantNode) secureURL() string {
	return tn.relativeSecureURL
}

func (tn *tenantNode) start(ctx context.Context, t test.Test, c cluster.Cluster, binary string) {
	require.True(t, c.IsSecure())

	tn.binary = binary
	extraArgs := []string{
		"--log=\"file-defaults: {dir: '" + tn.logDir() + "', exit-on-error: false}\"",
		"--store=" + tn.storeDir()}

	internalIPs, err := c.InternalIP(ctx, t.L(), c.Node(tn.node))
	require.NoError(t, err)
	tn.errCh = startTenantServer(
		ctx, c, c.Node(tn.node), internalIPs[0], binary, tn.kvAddrs, tn.tenantID,
		tn.httpPort, tn.sqlPort,
		extraArgs...,
	)

	externalUrls, err := c.ExternalPGUrl(ctx, t.L(), c.Node(tn.node))
	require.NoError(t, err)
	u, err := url.Parse(externalUrls[0])
	require.NoError(t, err)
	externalIPs, err := c.ExternalIP(ctx, t.L(), c.Node(tn.node))
	require.NoError(t, err)
	u.Host = externalIPs[0] + ":" + strconv.Itoa(tn.sqlPort)

	tn.pgURL = u.String()

	// pgURL has full paths to local certs embedded, i.e.
	// /tmp/roachtest-certs3630333874/certs, on the cluster we want just certs
	// (i.e. to run workload on the tenant).
	secureUrls, err := roachprod.PgURL(ctx, t.L(), c.MakeNodes(c.Node(tn.node)), "certs", false /*external*/, true /* secure */)
	require.NoError(t, err)
	u, err = url.Parse(strings.Trim(secureUrls[0], "'"))
	require.NoError(t, err)
	u.Host = externalIPs[0] + ":" + strconv.Itoa(tn.sqlPort)

	tn.relativeSecureURL = u.String()

	// The tenant is usually responsive ~right away, but it has on occasions
	// taken more than 3s for it to connect to the KV layer, and it won't open
	// the SQL port until it has.
	testutils.SucceedsSoon(t, func() error {
		select {
		case <-ctx.Done():
			t.Fatal(ctx.Err())
		case err := <-tn.errCh:
			t.Fatal(err)
		default:
		}

		db, err := gosql.Open("postgres", tn.pgURL)
		if err != nil {
			return err
		}
		defer db.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
		defer cancel()
		_, err = db.ExecContext(ctx, `SELECT 1`)
		return err
	})

	t.L().Printf("sql server for tenant %d (%d) running at %s", tn.tenantID, tn.instanceID, tn.pgURL)
}

func startTenantServer(
	tenantCtx context.Context,
	c cluster.Cluster,
	node option.NodeListOption,
	internalIP string,
	binary string,
	kvAddrs []string,
	tenantID int,
	httpPort int,
	sqlPort int,
	extraFlags ...string,
) chan error {
	args := []string{
		"--certs-dir", "certs",
		"--tenant-id=" + strconv.Itoa(tenantID),
		"--http-addr", ifLocal(c, "127.0.0.1", "0.0.0.0") + ":" + strconv.Itoa(httpPort),
		"--kv-addrs", strings.Join(kvAddrs, ","),
		// Don't bind to external interfaces when running locally.
		"--sql-addr", ifLocal(c, "127.0.0.1", internalIP) + ":" + strconv.Itoa(sqlPort),
	}
	args = append(args, extraFlags...)

	errCh := make(chan error, 1)
	go func() {
		errCh <- c.RunE(tenantCtx, node,
			append(append(append([]string{}, config.DefaultEnvVars()...), binary, "mt", "start-sql"), args...)...,
		)
		close(errCh)
	}()
	return errCh
}

func newTenantInstance(
	ctx context.Context, tn *tenantNode, t test.Test, c cluster.Cluster, node, http, sql int,
) (*tenantNode, error) {
	instID := tenantIds[tn.tenantID] + 1
	tenantIds[tn.tenantID] = instID
	inst := tenantNode{
		tenantID:   tn.tenantID,
		instanceID: instID,
		kvAddrs:    tn.kvAddrs,
		node:       node,
		httpPort:   http,
		sqlPort:    sql,
	}
	tenantCertsDir, err := os.MkdirTemp("", "tenant-certs")
	if err != nil {
		return nil, err
	}
	key, crt := fmt.Sprintf("client-tenant.%d.key", tn.tenantID), fmt.Sprintf("client-tenant.%d.crt", tn.tenantID)
	err = c.Get(ctx, t.L(), filepath.Join("certs", key), filepath.Join(tenantCertsDir, key), c.Node(tn.node))
	if err != nil {
		return nil, err
	}
	err = c.Get(ctx, t.L(), filepath.Join("certs", crt), filepath.Join(tenantCertsDir, crt), c.Node(tn.node))
	if err != nil {
		return nil, err
	}
	c.Put(ctx, filepath.Join(tenantCertsDir, key), filepath.Join("certs", key), c.Node(node))
	c.Put(ctx, filepath.Join(tenantCertsDir, crt), filepath.Join("certs", crt), c.Node(node))
	// sigh: locally theses are symlinked which breaks our crypto cert checks
	if c.IsLocal() {
		c.Run(ctx, c.Node(node), "rm", filepath.Join("certs", key))
		c.Run(ctx, c.Node(node), "cp", filepath.Join(tenantCertsDir, key), filepath.Join("certs", key))
	}
	c.Run(ctx, c.Node(node), "chmod", "0600", filepath.Join("certs", key))
	return &inst, nil
}
