// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Chassis(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Chassis error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Chassis result count: %d", len(result))
}
func TestGetBareMetal2Chassis(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2Chassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Chassis found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2Chassis(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2Chassis error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2Chassis result: %s", result.UUID)
}

func TestUpdateBareMetal2Chassis(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Chassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Chassis found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2ChassisParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2ChassisParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2Chassis(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Chassis error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2Chassis result: %s", result.UUID)
}

func TestDeleteBareMetal2Chassis(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBareMetal2Chassis is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2Chassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Chassis found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBareMetal2Chassis(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2Chassis error: %v", err)
		return
	}
	golog.Infof("DeleteBareMetal2Chassis succeeded for UUID: %s", list[0].UUID)
}

func TestInspectBareMetal2Chassis(t *testing.T) {
	// InspectBareMetal2Chassis operation
	t.Skip("TestInspectBareMetal2Chassis requires manual implementation")

}
