// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFlowMeter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFlowMeter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFlowMeter error: %v", err)
		return
	}
	golog.Infof("QueryFlowMeter result count: %d", len(result))
}

func TestUpdateFlowMeter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFlowMeter(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateFlowMeter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FlowMeter found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateFlowMeterParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateFlowMeterParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateFlowMeter(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateFlowMeter error: %v", err)
		return
	}
	golog.Infof("UpdateFlowMeter result: %s", result.UUID)
}

func TestDeleteFlowMeter(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteFlowMeter is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFlowMeter(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteFlowMeter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FlowMeter found to test Delete")
		return
	}

	err = accountLoginCli.DeleteFlowMeter(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteFlowMeter error: %v", err)
		return
	}
	golog.Infof("DeleteFlowMeter succeeded for UUID: %s", list[0].UUID)
}

func TestCreateFlowMeter(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateFlowMeter is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateFlowMeterParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateFlowMeterParamDetail{
	// 		Name: "test-flowmeter",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateFlowMeter(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateFlowMeter error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateFlowMeter result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteFlowMeter(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteFlowMeter error: %v", err)
	// }
}
