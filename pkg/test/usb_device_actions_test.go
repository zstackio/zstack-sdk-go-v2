// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUsbDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUsbDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUsbDevice error: %v", err)
		return
	}
	golog.Infof("QueryUsbDevice result count: %d", len(result))
}
func TestGetUsbDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUsbDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetUsbDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UsbDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetUsbDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUsbDevice error: %v", err)
		return
	}
	golog.Infof("GetUsbDevice result: %s", result.UUID)
}

func TestUpdateUsbDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUsbDevice(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateUsbDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UsbDevice found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateUsbDeviceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateUsbDeviceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateUsbDevice(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateUsbDevice error: %v", err)
		return
	}
	golog.Infof("UpdateUsbDevice result: %s", result.UUID)
}
