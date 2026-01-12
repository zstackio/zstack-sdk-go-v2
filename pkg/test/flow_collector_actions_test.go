// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFlowCollector(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFlowCollector(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFlowCollector error: %v", err)
		return
	}
	golog.Infof("QueryFlowCollector result count: %d", len(result))
}
func TestGetFlowCollector(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFlowCollector(&queryParam)
	if err != nil {
		t.Errorf("TestGetFlowCollector Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FlowCollector found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetFlowCollector(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetFlowCollector error: %v", err)
		return
	}
	golog.Infof("GetFlowCollector result: %s", result.UUID)
}

func TestUpdateFlowCollector(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFlowCollector(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateFlowCollector Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FlowCollector found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateFlowCollectorParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateFlowCollectorParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateFlowCollector(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateFlowCollector error: %v", err)
		return
	}
	golog.Infof("UpdateFlowCollector result: %s", result.UUID)
}

func TestDeleteFlowCollector(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteFlowCollector is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFlowCollector(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteFlowCollector Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FlowCollector found to test Delete")
		return
	}

	err = accountLoginCli.DeleteFlowCollector(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteFlowCollector error: %v", err)
		return
	}
	golog.Infof("DeleteFlowCollector succeeded for UUID: %s", list[0].UUID)
}

func TestCreateFlowCollector(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateFlowCollector is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateFlowCollectorParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateFlowCollectorParamDetail{
	// 		Name: "test-flowcollector",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateFlowCollector(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateFlowCollector error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateFlowCollector result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteFlowCollector(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteFlowCollector error: %v", err)
	// }
}
