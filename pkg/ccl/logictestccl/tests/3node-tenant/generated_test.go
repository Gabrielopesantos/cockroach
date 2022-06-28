// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by generate-logictest, DO NOT EDIT.

package test3node_tenant

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	_ "github.com/cockroachdb/cockroach/pkg/ccl"
	"github.com/cockroachdb/cockroach/pkg/ccl/utilccl"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 9

var logicTestDir string
var cclLogicTestDir string
var execBuildLogicTestDir string

func init() {
	if bazel.BuiltWithBazel() {
		var err error
		logicTestDir, err = bazel.Runfile("pkg/sql/logictest/testdata/logic_test")
		if err != nil {
			panic(err)
		}
	} else {
		logicTestDir = "../../../../sql/logictest/testdata/logic_test"
	}
	if bazel.BuiltWithBazel() {
		var err error
		cclLogicTestDir, err = bazel.Runfile("pkg/ccl/logictestccl/testdata/logic_test")
		if err != nil {
			panic(err)
		}
	} else {
		cclLogicTestDir = "../../../../ccl/logictestccl/testdata/logic_test"
	}
	if bazel.BuiltWithBazel() {
		var err error
		execBuildLogicTestDir, err = bazel.Runfile("pkg/sql/opt/exec/execbuilder/testdata")
		if err != nil {
			panic(err)
		}
	} else {
		execBuildLogicTestDir = "../../../../sql/opt/exec/execbuilder/testdata"
	}
}

func TestMain(m *testing.M) {
	defer utilccl.TestingEnableEnterprise()()
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}

func runLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	logictest.RunLogicTest(t, logictest.TestServerArgs{}, configIdx, filepath.Join(logicTestDir, file))
}
func runCCLLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	logictest.RunLogicTest(t, logictest.TestServerArgs{}, configIdx, filepath.Join(cclLogicTestDir, file))
}
func runExecBuildLogicTest(t *testing.T, file string) {
	defer sql.TestingOverrideExplainEnvVersion("CockroachDB execbuilder test version")()
	skip.UnderDeadlock(t, "times out and/or hangs")
	serverArgs := logictest.TestServerArgs{
		DisableWorkmemRandomization: true,
	}
	logictest.RunLogicTest(t, serverArgs, configIdx, filepath.Join(execBuildLogicTestDir, file))
}

func TestTenantLogic_aggregate(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "aggregate")
}

func TestTenantLogic_alias_types(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alias_types")
}

func TestTenantLogic_alter_column_type(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_column_type")
}

func TestTenantLogic_alter_database_convert_to_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_database_convert_to_schema")
}

func TestTenantLogic_alter_database_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_database_owner")
}

func TestTenantLogic_alter_default_privileges_for_all_roles(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_for_all_roles")
}

func TestTenantLogic_alter_default_privileges_for_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_for_schema")
}

func TestTenantLogic_alter_default_privileges_for_sequence(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_for_sequence")
}

func TestTenantLogic_alter_default_privileges_for_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_for_table")
}

func TestTenantLogic_alter_default_privileges_for_type(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_for_type")
}

func TestTenantLogic_alter_default_privileges_in_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_in_schema")
}

func TestTenantLogic_alter_default_privileges_with_grant_option(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_default_privileges_with_grant_option")
}

func TestTenantLogic_alter_primary_key(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_primary_key")
}

func TestTenantLogic_alter_role(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_role")
}

func TestTenantLogic_alter_role_set(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_role_set")
}

func TestTenantLogic_alter_schema_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_schema_owner")
}

func TestTenantLogic_alter_sequence(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_sequence")
}

func TestTenantLogic_alter_sequence_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_sequence_owner")
}

func TestTenantLogic_alter_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_table")
}

func TestTenantLogic_alter_table_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_table_owner")
}

func TestTenantLogic_alter_type(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_type")
}

func TestTenantLogic_alter_type_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_type_owner")
}

func TestTenantLogic_alter_view_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "alter_view_owner")
}

func TestTenantLogic_and_or(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "and_or")
}

func TestTenantLogic_apply_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "apply_join")
}

func TestTenantLogic_array(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "array")
}

func TestTenantLogic_as_of(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "as_of")
}

func TestTenantLogic_auto_span_config_reconciliation_job(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "auto_span_config_reconciliation_job")
}

func TestTenantLogic_bit(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "bit")
}

func TestTenantLogic_builtin_function(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "builtin_function")
}

func TestTenantLogic_bytes(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "bytes")
}

func TestTenantLogic_cascade(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "cascade")
}

func TestTenantLogic_case_sensitive_names(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "case_sensitive_names")
}

func TestTenantLogic_cast(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "cast")
}

func TestTenantLogic_check_constraints(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "check_constraints")
}

func TestTenantLogic_cluster_settings(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "cluster_settings")
}

func TestTenantLogic_collatedstring(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring")
}

func TestTenantLogic_collatedstring_constraint(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_constraint")
}

func TestTenantLogic_collatedstring_index1(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_index1")
}

func TestTenantLogic_collatedstring_index2(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_index2")
}

func TestTenantLogic_collatedstring_normalization(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_normalization")
}

func TestTenantLogic_collatedstring_nullinindex(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_nullinindex")
}

func TestTenantLogic_collatedstring_uniqueindex1(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_uniqueindex1")
}

func TestTenantLogic_collatedstring_uniqueindex2(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "collatedstring_uniqueindex2")
}

func TestTenantLogic_column_families(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "column_families")
}

func TestTenantLogic_computed(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "computed")
}

func TestTenantLogic_conditional(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "conditional")
}

func TestTenantLogic_connect_privilege(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "connect_privilege")
}

func TestTenantLogic_crdb_internal_default_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "crdb_internal_default_privileges")
}

func TestTenantLogic_create_as(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "create_as")
}

func TestTenantLogic_create_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "create_index")
}

func TestTenantLogic_create_statements(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "create_statements")
}

func TestTenantLogic_create_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "create_table")
}

func TestTenantLogic_cross_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "cross_join")
}

func TestTenantLogic_cursor(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "cursor")
}

func TestTenantLogic_custom_escape_character(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "custom_escape_character")
}

func TestTenantLogic_dangerous_statements(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "dangerous_statements")
}

func TestTenantLogic_database(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "database")
}

func TestTenantLogic_datetime(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "datetime")
}

func TestTenantLogic_decimal(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "decimal")
}

func TestTenantLogic_default(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "default")
}

func TestTenantLogic_delete(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "delete")
}

func TestTenantLogic_dependencies(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "dependencies")
}

func TestTenantLogic_discard(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "discard")
}

func TestTenantLogic_disjunction_in_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "disjunction_in_join")
}

func TestTenantLogic_distinct(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distinct")
}

func TestTenantLogic_distinct_on(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distinct_on")
}

func TestTenantLogic_distsql_automatic_stats(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_automatic_stats")
}

func TestTenantLogic_distsql_event_log(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_event_log")
}

func TestTenantLogic_distsql_expr(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_expr")
}

func TestTenantLogic_distsql_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_join")
}

func TestTenantLogic_distsql_srfs(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_srfs")
}

func TestTenantLogic_distsql_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "distsql_tenant")
}

func TestTenantLogic_drop_database(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_database")
}

func TestTenantLogic_drop_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_index")
}

func TestTenantLogic_drop_owned_by(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_owned_by")
}

func TestTenantLogic_drop_role_with_default_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_role_with_default_privileges")
}

func TestTenantLogic_drop_role_with_default_privileges_in_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_role_with_default_privileges_in_schema")
}

func TestTenantLogic_drop_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_schema")
}

func TestTenantLogic_drop_sequence(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_sequence")
}

func TestTenantLogic_drop_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_table")
}

func TestTenantLogic_drop_type(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_type")
}

func TestTenantLogic_drop_user(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_user")
}

func TestTenantLogic_drop_view(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "drop_view")
}

func TestTenantLogic_edge(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "edge")
}

func TestTenantLogic_enums(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "enums")
}

func TestTenantLogic_errors(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "errors")
}

func TestTenantLogic_event_log(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "event_log")
}

func TestTenantLogic_exclude_data_from_backup(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "exclude_data_from_backup")
}

func TestTenantLogic_experimental_distsql_planning(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "experimental_distsql_planning")
}

func TestTenantLogic_explain_analyze(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "explain_analyze")
}

func TestTenantLogic_expression_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "expression_index")
}

func TestTenantLogic_family(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "family")
}

func TestTenantLogic_fk(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "fk")
}

func TestTenantLogic_float(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "float")
}

func TestTenantLogic_function_lookup(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "function_lookup")
}

func TestTenantLogic_fuzzystrmatch(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "fuzzystrmatch")
}

func TestTenantLogic_geospatial(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "geospatial")
}

func TestTenantLogic_geospatial_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "geospatial_index")
}

func TestTenantLogic_geospatial_meta(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "geospatial_meta")
}

func TestTenantLogic_geospatial_zm(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "geospatial_zm")
}

func TestTenantLogic_grant_database(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_database")
}

func TestTenantLogic_grant_in_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_in_txn")
}

func TestTenantLogic_grant_on_all_sequences_in_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_on_all_sequences_in_schema")
}

func TestTenantLogic_grant_on_all_tables_in_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_on_all_tables_in_schema")
}

func TestTenantLogic_grant_revoke_with_grant_option(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_revoke_with_grant_option")
}

func TestTenantLogic_grant_role(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_role")
}

func TestTenantLogic_grant_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "grant_schema")
}

func TestTenantLogic_hash_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "hash_join")
}

func TestTenantLogic_hash_sharded_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "hash_sharded_index")
}

func TestTenantLogic_hidden_columns(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "hidden_columns")
}

func TestTenantLogic_impure(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "impure")
}

func TestTenantLogic_index_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "index_join")
}

func TestTenantLogic_inet(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inet")
}

func TestTenantLogic_inflight_trace_spans(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inflight_trace_spans")
}

func TestTenantLogic_inner_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inner-join")
}

func TestTenantLogic_insert(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "insert")
}

func TestTenantLogic_int_size(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "int_size")
}

func TestTenantLogic_internal_executor(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "internal_executor")
}

func TestTenantLogic_interval(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "interval")
}

func TestTenantLogic_inv_stats(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inv_stats")
}

func TestTenantLogic_inverted_filter_geospatial(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_filter_geospatial")
}

func TestTenantLogic_inverted_filter_json_array(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_filter_json_array")
}

func TestTenantLogic_inverted_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_index")
}

func TestTenantLogic_inverted_index_multi_column(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_index_multi_column")
}

func TestTenantLogic_inverted_join_geospatial(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_join_geospatial")
}

func TestTenantLogic_inverted_join_json_array(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_join_json_array")
}

func TestTenantLogic_inverted_join_multi_column(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "inverted_join_multi_column")
}

func TestTenantLogic_jobs(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "jobs")
}

func TestTenantLogic_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "join")
}

func TestTenantLogic_json(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "json")
}

func TestTenantLogic_json_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "json_builtins")
}

func TestTenantLogic_kv_builtin_functions_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "kv_builtin_functions_tenant")
}

func TestTenantLogic_limit(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "limit")
}

func TestTenantLogic_lock_timeout(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "lock_timeout")
}

func TestTenantLogic_lookup_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "lookup_join")
}

func TestTenantLogic_lookup_join_spans(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "lookup_join_spans")
}

func TestTenantLogic_manual_retry(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "manual_retry")
}

func TestTenantLogic_materialized_view(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "materialized_view")
}

func TestTenantLogic_merge_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "merge_join")
}

func TestTenantLogic_multi_statement(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "multi_statement")
}

func TestTenantLogic_name_escapes(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "name_escapes")
}

func TestTenantLogic_namespace(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "namespace")
}

func TestTenantLogic_new_schema_changer(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "new_schema_changer")
}

func TestTenantLogic_no_primary_key(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "no_primary_key")
}

func TestTenantLogic_notice(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "notice")
}

func TestTenantLogic_numeric_references(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "numeric_references")
}

func TestTenantLogic_on_update(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "on_update")
}

func TestTenantLogic_operator(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "operator")
}

func TestTenantLogic_optimizer_timeout(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "optimizer_timeout")
}

func TestTenantLogic_order_by(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "order_by")
}

func TestTenantLogic_ordinal_references(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "ordinal_references")
}

func TestTenantLogic_ordinality(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "ordinality")
}

func TestTenantLogic_orms(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "orms")
}

func TestTenantLogic_overflow(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "overflow")
}

func TestTenantLogic_overlaps(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "overlaps")
}

func TestTenantLogic_owner(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "owner")
}

func TestTenantLogic_parallel_stmts_compat(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "parallel_stmts_compat")
}

func TestTenantLogic_partial_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "partial_index")
}

func TestTenantLogic_partial_txn_commit(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "partial_txn_commit")
}

func TestTenantLogic_pg_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pg_builtins")
}

func TestTenantLogic_pg_catalog_pg_default_acl(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pg_catalog_pg_default_acl")
}

func TestTenantLogic_pg_catalog_pg_default_acl_with_grant_option(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pg_catalog_pg_default_acl_with_grant_option")
}

func TestTenantLogic_pg_extension(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pg_extension")
}

func TestTenantLogic_pgcrypto_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pgcrypto_builtins")
}

func TestTenantLogic_pgoidtype(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "pgoidtype")
}

func TestTenantLogic_poison_after_push(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "poison_after_push")
}

func TestTenantLogic_postgres_jsonb(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "postgres_jsonb")
}

func TestTenantLogic_postgresjoin(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "postgresjoin")
}

func TestTenantLogic_privilege_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "privilege_builtins")
}

func TestTenantLogic_privileges_comments(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "privileges_comments")
}

func TestTenantLogic_privileges_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "privileges_table")
}

func TestTenantLogic_propagate_input_ordering(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "propagate_input_ordering")
}

func TestTenantLogic_reassign_owned_by(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "reassign_owned_by")
}

func TestTenantLogic_record(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "record")
}

func TestTenantLogic_rename_atomic(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_atomic")
}

func TestTenantLogic_rename_column(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_column")
}

func TestTenantLogic_rename_constraint(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_constraint")
}

func TestTenantLogic_rename_database(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_database")
}

func TestTenantLogic_rename_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_index")
}

func TestTenantLogic_rename_sequence(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_sequence")
}

func TestTenantLogic_rename_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_table")
}

func TestTenantLogic_rename_view(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rename_view")
}

func TestTenantLogic_reset(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "reset")
}

func TestTenantLogic_returning(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "returning")
}

func TestTenantLogic_role(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "role")
}

func TestTenantLogic_row_level_ttl(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "row_level_ttl")
}

func TestTenantLogic_rows_from(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "rows_from")
}

func TestTenantLogic_run_control(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "run_control")
}

func TestTenantLogic_save_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "save_table")
}

func TestTenantLogic_scale(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "scale")
}

func TestTenantLogic_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "schema")
}

func TestTenantLogic_schema_change_feature_flags(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "schema_change_feature_flags")
}

func TestTenantLogic_schema_change_in_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "schema_change_in_txn")
}

func TestTenantLogic_schema_change_retry(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "schema_change_retry")
}

func TestTenantLogic_schema_repair(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "schema_repair")
}

func TestTenantLogic_scrub(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "scrub")
}

func TestTenantLogic_secondary_index_column_families(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "secondary_index_column_families")
}

func TestTenantLogic_select(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select")
}

func TestTenantLogic_select_for_update(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select_for_update")
}

func TestTenantLogic_select_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select_index")
}

func TestTenantLogic_select_index_flags(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select_index_flags")
}

func TestTenantLogic_select_search_path(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select_search_path")
}

func TestTenantLogic_select_table_alias(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "select_table_alias")
}

func TestTenantLogic_sequences(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "sequences")
}

func TestTenantLogic_sequences_distsql(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "sequences_distsql")
}

func TestTenantLogic_sequences_regclass(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "sequences_regclass")
}

func TestTenantLogic_serial(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "serial")
}

func TestTenantLogic_serializable_eager_restart(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "serializable_eager_restart")
}

func TestTenantLogic_set_local(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "set_local")
}

func TestTenantLogic_set_role(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "set_role")
}

func TestTenantLogic_set_schema(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "set_schema")
}

func TestTenantLogic_set_time_zone(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "set_time_zone")
}

func TestTenantLogic_shift(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "shift")
}

func TestTenantLogic_show_completions(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_completions")
}

func TestTenantLogic_show_create(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_create")
}

func TestTenantLogic_show_create_all_schemas(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_create_all_schemas")
}

func TestTenantLogic_show_create_all_tables(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_create_all_tables")
}

func TestTenantLogic_show_create_all_tables_builtin(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_create_all_tables_builtin")
}

func TestTenantLogic_show_create_all_types(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_create_all_types")
}

func TestTenantLogic_show_default_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_default_privileges")
}

func TestTenantLogic_show_fingerprints(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_fingerprints")
}

func TestTenantLogic_show_indexes(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "show_indexes")
}

func TestTenantLogic_span_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "span_builtins")
}

func TestTenantLogic_sqllite(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "sqllite")
}

func TestTenantLogic_sqlsmith(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "sqlsmith")
}

func TestTenantLogic_srfs(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "srfs")
}

func TestTenantLogic_statement_source(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "statement_source")
}

func TestTenantLogic_statement_statistics(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "statement_statistics")
}

func TestTenantLogic_storing(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "storing")
}

func TestTenantLogic_strict_ddl_atomicity(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "strict_ddl_atomicity")
}

func TestTenantLogic_suboperators(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "suboperators")
}

func TestTenantLogic_subquery(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "subquery")
}

func TestTenantLogic_subquery_correlated(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "subquery_correlated")
}

func TestTenantLogic_system(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "system")
}

func TestTenantLogic_system_columns(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "system_columns")
}

func TestTenantLogic_system_namespace(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "system_namespace")
}

func TestTenantLogic_system_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "system_privileges")
}

func TestTenantLogic_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "table")
}

func TestTenantLogic_target_names(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "target_names")
}

func TestTenantLogic_temp_table(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "temp_table")
}

func TestTenantLogic_temp_table_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "temp_table_txn")
}

func TestTenantLogic_tenant_slow_repro(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "tenant_slow_repro")
}

func TestTenantLogic_time(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "time")
}

func TestTenantLogic_timestamp(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "timestamp")
}

func TestTenantLogic_timetz(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "timetz")
}

func TestTenantLogic_trigram_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "trigram_builtins")
}

func TestTenantLogic_trigram_indexes(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "trigram_indexes")
}

func TestTenantLogic_truncate(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "truncate")
}

func TestTenantLogic_tuple(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "tuple")
}

func TestTenantLogic_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "txn")
}

func TestTenantLogic_txn_as_of(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "txn_as_of")
}

func TestTenantLogic_txn_retry(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "txn_retry")
}

func TestTenantLogic_txn_stats(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "txn_stats")
}

func TestTenantLogic_type_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "type_privileges")
}

func TestTenantLogic_typing(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "typing")
}

func TestTenantLogic_union(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "union")
}

func TestTenantLogic_unique(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "unique")
}

func TestTenantLogic_update(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "update")
}

func TestTenantLogic_update_from(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "update_from")
}

func TestTenantLogic_upsert(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "upsert")
}

func TestTenantLogic_user(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "user")
}

func TestTenantLogic_uuid(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "uuid")
}

func TestTenantLogic_values(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "values")
}

func TestTenantLogic_vectorize(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize")
}

func TestTenantLogic_vectorize_agg(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_agg")
}

func TestTenantLogic_vectorize_overloads(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_overloads")
}

func TestTenantLogic_vectorize_shutdown(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_shutdown")
}

func TestTenantLogic_vectorize_types(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_types")
}

func TestTenantLogic_vectorize_unsupported(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_unsupported")
}

func TestTenantLogic_vectorize_window(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "vectorize_window")
}

func TestTenantLogic_views(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "views")
}

func TestTenantLogic_virtual_columns(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "virtual_columns")
}

func TestTenantLogic_virtual_table_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "virtual_table_privileges")
}

func TestTenantLogic_void(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "void")
}

func TestTenantLogic_where(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "where")
}

func TestTenantLogic_window(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "window")
}

func TestTenantLogic_with(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "with")
}

func TestTenantLogic_zero(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "zero")
}

func TestTenantLogic_zigzag_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "zigzag_join")
}

func TestTenantLogic_zone_config(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "zone_config")
}

func TestTenantLogicCCL_cluster_locks_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "cluster_locks_tenant")
}

func TestTenantLogicCCL_crdb_internal_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "crdb_internal_tenant")
}

func TestTenantLogicCCL_new_schema_changer(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "new_schema_changer")
}

func TestTenantLogicCCL_partitioning_enum(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_enum")
}

func TestTenantLogicCCL_schema_change_in_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "schema_change_in_txn")
}

func TestTenantLogicCCL_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "tenant")
}

func TestTenantLogicCCL_tenant_settings(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "tenant_settings")
}

func TestTenantLogicCCL_tenant_unsupported(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "tenant_unsupported")
}

func TestTenantLogicCCL_zone_config_secondary_tenants(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "zone_config_secondary_tenants")
}

func TestTenantLogicCCL_zone_config_secondary_tenants_disallowed(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "zone_config_secondary_tenants_disallowed")
}

func TestTenantExecBuild_distsql_tenant(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_tenant")
}
