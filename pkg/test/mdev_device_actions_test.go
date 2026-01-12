// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMdevDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMdevDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMdevDevice error: %v", err)
		return
	}
	golog.Infof("QueryMdevDevice result count: %d", len(result))
}
func TestGetMdevDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMdevDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetMdevDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MdevDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMdevDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMdevDevice error: %v", err)
		return
	}
	golog.Infof("GetMdevDevice result: %s", result.UUID)
}

func TestUpdateMdevDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMdevDevice(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMdevDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MdevDevice found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMdevDeviceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMdevDeviceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMdevDevice(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMdevDevice error: %v", err)
		return
	}
	golog.Infof("UpdateMdevDevice result: %s", result.UUID)
}

func TestDeleteMdevDevice(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMdevDevice is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMdevDevice(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMdevDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MdevDevice found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMdevDevice(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMdevDevice error: %v", err)
		return
	}
	golog.Infof("DeleteMdevDevice succeeded for UUID: %s", list[0].UUID)
}
