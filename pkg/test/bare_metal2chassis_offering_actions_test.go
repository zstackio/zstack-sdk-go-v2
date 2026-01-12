// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ChassisOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ChassisOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ChassisOffering error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ChassisOffering result count: %d", len(result))
}
func TestGetBareMetal2ChassisOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ChassisOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ChassisOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2ChassisOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisOffering error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2ChassisOffering result: %s", result.UUID)
}

func TestUpdateBareMetal2ChassisOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ChassisOffering(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ChassisOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ChassisOffering found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2ChassisOfferingParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2ChassisOfferingParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2ChassisOffering(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ChassisOffering error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2ChassisOffering result: %s", result.UUID)
}
