// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBuildAppExportHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBuildAppExportHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBuildAppExportHistory error: %v", err)
		return
	}
	golog.Infof("QueryBuildAppExportHistory result count: %d", len(result))
}
func TestGetBuildAppExportHistory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBuildAppExportHistory(&queryParam)
	if err != nil {
		t.Errorf("TestGetBuildAppExportHistory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BuildAppExportHistory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBuildAppExportHistory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBuildAppExportHistory error: %v", err)
		return
	}
	golog.Infof("GetBuildAppExportHistory result: %s", result.UUID)
}

func TestDeleteBuildAppExportHistory(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBuildAppExportHistory is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBuildAppExportHistory(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBuildAppExportHistory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BuildAppExportHistory found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBuildAppExportHistory(*list[0].BuildAppUuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBuildAppExportHistory error: %v", err)
		return
	}
	golog.Infof("DeleteBuildAppExportHistory succeeded for UUID: %s", list[0].BuildAppUuid)
}
