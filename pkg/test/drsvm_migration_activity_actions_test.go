// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDRSVmMigrationActivity(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDRSVmMigrationActivity(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDRSVmMigrationActivity error: %v", err)
		return
	}
	golog.Infof("QueryDRSVmMigrationActivity result count: %d", len(result))
}
func TestGetDRSVmMigrationActivity(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDRSVmMigrationActivity(&queryParam)
	if err != nil {
		t.Errorf("TestGetDRSVmMigrationActivity Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DRSVmMigrationActivity found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetDRSVmMigrationActivity(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDRSVmMigrationActivity error: %v", err)
		return
	}
	golog.Infof("GetDRSVmMigrationActivity result: %s", result.UUID)
}
