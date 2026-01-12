// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenter error: %v", err)
		return
	}
	golog.Infof("QueryVCenter result count: %d", len(result))
}
func TestGetVCenter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenter(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenter found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenter(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenter error: %v", err)
		return
	}
	golog.Infof("GetVCenter result: %s", result.UUID)
}

func TestUpdateVCenter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenter(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVCenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenter found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVCenterParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVCenterParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVCenter(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVCenter error: %v", err)
		return
	}
	golog.Infof("UpdateVCenter result: %s", result.UUID)
}

func TestDeleteVCenter(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVCenter is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenter(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVCenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenter found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVCenter(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVCenter error: %v", err)
		return
	}
	golog.Infof("DeleteVCenter succeeded for UUID: %s", list[0].UUID)
}

func TestAddVCenter(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddVCenter requires valid creation parameters")

}

func TestSyncVCenter(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncVCenter requires a valid resource to sync")

}
