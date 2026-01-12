// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPciDeviceSpec(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPciDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDeviceSpec error: %v", err)
		return
	}
	golog.Infof("QueryPciDeviceSpec result count: %d", len(result))
}

func TestUpdatePciDeviceSpec(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPciDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePciDeviceSpec Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDeviceSpec found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePciDeviceSpecParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePciDeviceSpecParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePciDeviceSpec(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePciDeviceSpec error: %v", err)
		return
	}
	golog.Infof("UpdatePciDeviceSpec result: %s", result.UUID)
}
