// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import { createMemoryHistory } from "history";
import _ from "lodash";
import Long from "long";
import { bindActionCreators, Store } from "redux";
import {
  DatabasesPageActions,
  DatabasesPageData,
  DatabasesPageDataDatabase,
  DatabasesPageDataMissingTable,
} from "@cockroachlabs/cluster-ui";

import { AdminUIState, createAdminUIStore } from "src/redux/state";
import * as fakeApi from "src/util/fakeApi";
import { mapDispatchToProps, mapStateToProps } from "./redux";

class TestDriver {
  private readonly actions: DatabasesPageActions;
  private readonly properties: () => DatabasesPageData;

  constructor(store: Store<AdminUIState>) {
    this.actions = bindActionCreators(
      mapDispatchToProps,
      store.dispatch.bind(store),
    );
    this.properties = () => mapStateToProps(store.getState());
  }

  async refreshDatabases() {
    return this.actions.refreshDatabases();
  }

  async refreshDatabaseDetails(database: string) {
    return this.actions.refreshDatabaseDetails(database);
  }

  async refreshTableStats(database: string, table: string) {
    return this.actions.refreshTableStats(database, table);
  }

  async refreshSettings() {
    return this.actions.refreshSettings();
  }

  assertProperties(expected: DatabasesPageData) {
    expect(this.properties()).toEqual(expected);
  }

  assertDatabaseProperties(
    database: string,
    expected: DatabasesPageDataDatabase,
  ) {
    expect(this.findDatabase(database)).toEqual(expected);
  }

  assertMissingTableProperties(
    database: string,
    table: string,
    expected: DatabasesPageDataMissingTable,
  ) {
    expect(this.findMissingTable(this.findDatabase(database), table)).toEqual(
      expected,
    );
  }

  private findDatabase(name: string) {
    return _.find(this.properties().databases, row => row.name == name);
  }

  private findMissingTable(database: DatabasesPageDataDatabase, name: string) {
    return _.find(database.missingTables, table => table.name == name);
  }
}

describe("Databases Page", function () {
  let driver: TestDriver;

  beforeEach(function () {
    driver = new TestDriver(createAdminUIStore(createMemoryHistory()));
  });

  afterEach(function () {
    fakeApi.restore();
  });

  it("starts in a pre-loading state", async function () {
    fakeApi.stubClusterSettings({
      key_values: {
        "sql.stats.automatic_collection.enabled": { value: "true" },
      },
    });

    await driver.refreshSettings();

    driver.assertProperties({
      loading: false,
      loaded: false,
      lastError: undefined,
      databases: [],
      sortSetting: { ascending: true, columnTitle: "name" },
      automaticStatsCollectionEnabled: true,
      showNodeRegionsColumn: false,
    });
  });

  it("makes a row for each database", async function () {
    fakeApi.stubDatabases(["system", "test"]);
    fakeApi.stubClusterSettings({
      key_values: {
        "sql.stats.automatic_collection.enabled": { value: "true" },
      },
    });

    await driver.refreshDatabases();
    await driver.refreshSettings();

    driver.assertProperties({
      loading: false,
      loaded: true,
      lastError: null,
      databases: [
        {
          loading: false,
          loaded: false,
          lastError: undefined,
          name: "system",
          sizeInBytes: 0,
          tableCount: 0,
          rangeCount: 0,
          nodesByRegionString: "",
          missingTables: [],
          numIndexRecommendations: 0,
        },
        {
          loading: false,
          loaded: false,
          lastError: undefined,
          name: "test",
          sizeInBytes: 0,
          tableCount: 0,
          rangeCount: 0,
          nodesByRegionString: "",
          missingTables: [],
          numIndexRecommendations: 0,
        },
      ],
      sortSetting: { ascending: true, columnTitle: "name" },
      showNodeRegionsColumn: false,
      automaticStatsCollectionEnabled: true,
    });
  });

  it("fills in database details", async function () {
    fakeApi.stubDatabases(["system", "test"]);

    fakeApi.stubDatabaseDetails("system", {
      table_names: ["foo", "bar"],
      stats: {
        missing_tables: [],
        range_count: new Long(3),
        approximate_disk_bytes: new Long(7168),
      },
    });

    fakeApi.stubDatabaseDetails("test", {
      table_names: ["widgets"],
      stats: {
        missing_tables: [],
        range_count: new Long(42),
        approximate_disk_bytes: new Long(1234),
      },
    });

    await driver.refreshDatabases();
    await driver.refreshDatabaseDetails("system");
    await driver.refreshDatabaseDetails("test");

    driver.assertDatabaseProperties("system", {
      loading: false,
      loaded: true,
      lastError: null,
      name: "system",
      sizeInBytes: 7168,
      tableCount: 2,
      rangeCount: 3,
      nodesByRegionString: "",
      missingTables: [],
      numIndexRecommendations: 0,
    });

    driver.assertDatabaseProperties("test", {
      loading: false,
      loaded: true,
      lastError: null,
      name: "test",
      sizeInBytes: 1234,
      tableCount: 1,
      rangeCount: 42,
      nodesByRegionString: "",
      missingTables: [],
      numIndexRecommendations: 0,
    });
  });

  describe("fallback cases", function () {
    describe("missing tables", function () {
      it("exposes them so the component can refresh them", async function () {
        fakeApi.stubDatabases(["system"]);

        fakeApi.stubDatabaseDetails("system", {
          table_names: ["foo", "bar"],
          stats: {
            missing_tables: [{ name: "bar" }],
            range_count: new Long(3),
            approximate_disk_bytes: new Long(7168),
          },
        });

        await driver.refreshDatabases();
        await driver.refreshDatabaseDetails("system");

        driver.assertDatabaseProperties("system", {
          loading: false,
          loaded: true,
          lastError: null,
          name: "system",
          sizeInBytes: 7168,
          tableCount: 2,
          rangeCount: 3,
          nodesByRegionString: "",
          missingTables: [{ loading: false, name: "bar" }],
          numIndexRecommendations: 0,
        });
      });

      it("merges available individual stats into the totals", async function () {
        fakeApi.stubDatabases(["system"]);

        fakeApi.stubDatabaseDetails("system", {
          table_names: ["foo", "bar"],
          stats: {
            missing_tables: [{ name: "bar" }],
            range_count: new Long(3),
            approximate_disk_bytes: new Long(7168),
          },
        });

        fakeApi.stubTableStats("system", "bar", {
          range_count: new Long(5),
          approximate_disk_bytes: new Long(1024),
        });

        await driver.refreshDatabases();
        await driver.refreshDatabaseDetails("system");
        await driver.refreshTableStats("system", "bar");

        driver.assertDatabaseProperties("system", {
          loading: false,
          loaded: true,
          lastError: null,
          name: "system",
          sizeInBytes: 8192,
          tableCount: 2,
          rangeCount: 8,
          nodesByRegionString: "",
          missingTables: [],
          numIndexRecommendations: 0,
        });
      });
    });

    describe("missing stats", function () {
      it("builds a list of missing tables", async function () {
        fakeApi.stubDatabases(["system"]);

        fakeApi.stubDatabaseDetails("system", {
          table_names: ["foo", "bar"],
        });

        await driver.refreshDatabases();
        await driver.refreshDatabaseDetails("system");

        driver.assertDatabaseProperties("system", {
          loading: false,
          loaded: true,
          lastError: null,
          name: "system",
          sizeInBytes: 0,
          tableCount: 2,
          rangeCount: 0,
          nodesByRegionString: "",
          missingTables: [
            { loading: false, name: "foo" },
            { loading: false, name: "bar" },
          ],
          numIndexRecommendations: 0,
        });
      });

      it("merges individual stats into the totals", async function () {
        fakeApi.stubDatabases(["system"]);

        fakeApi.stubDatabaseDetails("system", {
          table_names: ["foo", "bar"],
        });

        fakeApi.stubTableStats("system", "foo", {
          range_count: new Long(3),
          approximate_disk_bytes: new Long(7168),
        });

        fakeApi.stubTableStats("system", "bar", {
          range_count: new Long(5),
          approximate_disk_bytes: new Long(1024),
        });

        await driver.refreshDatabases();
        await driver.refreshDatabaseDetails("system");
        await driver.refreshTableStats("system", "foo");

        driver.assertDatabaseProperties("system", {
          loading: false,
          loaded: true,
          lastError: null,
          name: "system",
          sizeInBytes: 7168,
          tableCount: 2,
          rangeCount: 3,
          nodesByRegionString: "",
          missingTables: [{ loading: false, name: "bar" }],
          numIndexRecommendations: 0,
        });

        await driver.refreshTableStats("system", "bar");

        driver.assertDatabaseProperties("system", {
          loading: false,
          loaded: true,
          lastError: null,
          name: "system",
          sizeInBytes: 8192,
          tableCount: 2,
          rangeCount: 8,
          nodesByRegionString: "",
          missingTables: [],
          numIndexRecommendations: 0,
        });
      });
    });
  });
});
