// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroup error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroup result count: %d", len(result))
}
func TestGetMonitorGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMonitorGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorGroup error: %v", err)
		return
	}
	golog.Infof("GetMonitorGroup result: %s", result.UUID)
}

func TestUpdateMonitorGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMonitorGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMonitorGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMonitorGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorGroup error: %v", err)
		return
	}
	golog.Infof("UpdateMonitorGroup result: %s", result.UUID)
}

func TestDeleteMonitorGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMonitorGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMonitorGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMonitorGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMonitorGroup error: %v", err)
		return
	}
	golog.Infof("DeleteMonitorGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMonitorGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMonitorGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMonitorGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMonitorGroupParamDetail{
	// 		Name: "test-monitorgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMonitorGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMonitorGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMonitorGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMonitorGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMonitorGroup error: %v", err)
	// }
}
