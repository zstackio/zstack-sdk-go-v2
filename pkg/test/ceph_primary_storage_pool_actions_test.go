// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStoragePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephPrimaryStoragePool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStoragePool error: %v", err)
		return
	}
	golog.Infof("QueryCephPrimaryStoragePool result count: %d", len(result))
}
func TestGetCephPrimaryStoragePool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephPrimaryStoragePool(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStoragePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStoragePool found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCephPrimaryStoragePool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStoragePool error: %v", err)
		return
	}
	golog.Infof("GetCephPrimaryStoragePool result: %s", result.UUID)
}

func TestUpdateCephPrimaryStoragePool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephPrimaryStoragePool(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCephPrimaryStoragePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStoragePool found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCephPrimaryStoragePoolParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCephPrimaryStoragePoolParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCephPrimaryStoragePool(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCephPrimaryStoragePool error: %v", err)
		return
	}
	golog.Infof("UpdateCephPrimaryStoragePool result: %s", result.UUID)
}

func TestDeleteCephPrimaryStoragePool(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCephPrimaryStoragePool is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephPrimaryStoragePool(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCephPrimaryStoragePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStoragePool found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCephPrimaryStoragePool(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCephPrimaryStoragePool error: %v", err)
		return
	}
	golog.Infof("DeleteCephPrimaryStoragePool succeeded for UUID: %s", list[0].UUID)
}

func TestAddCephPrimaryStoragePool(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCephPrimaryStoragePool requires valid creation parameters")

}
