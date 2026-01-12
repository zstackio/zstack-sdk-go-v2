// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMdevDeviceSpec(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMdevDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMdevDeviceSpec error: %v", err)
		return
	}
	golog.Infof("QueryMdevDeviceSpec result count: %d", len(result))
}
func TestGetMdevDeviceSpec(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMdevDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestGetMdevDeviceSpec Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MdevDeviceSpec found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMdevDeviceSpec(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMdevDeviceSpec error: %v", err)
		return
	}
	golog.Infof("GetMdevDeviceSpec result: %s", result.UUID)
}

func TestUpdateMdevDeviceSpec(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMdevDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMdevDeviceSpec Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MdevDeviceSpec found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMdevDeviceSpecParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMdevDeviceSpecParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMdevDeviceSpec(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMdevDeviceSpec error: %v", err)
		return
	}
	golog.Infof("UpdateMdevDeviceSpec result: %s", result.UUID)
}
