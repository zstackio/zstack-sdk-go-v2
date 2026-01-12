// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ChassisPciDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ChassisPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ChassisPciDevice error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ChassisPciDevice result count: %d", len(result))
}
func TestGetBareMetal2ChassisPciDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ChassisPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisPciDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ChassisPciDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2ChassisPciDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisPciDevice error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2ChassisPciDevice result: %s", result.UUID)
}

func TestUpdateBareMetal2ChassisPciDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ChassisPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ChassisPciDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ChassisPciDevice found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2ChassisPciDeviceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2ChassisPciDeviceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2ChassisPciDevice(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ChassisPciDevice error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2ChassisPciDevice result: %s", result.UUID)
}
