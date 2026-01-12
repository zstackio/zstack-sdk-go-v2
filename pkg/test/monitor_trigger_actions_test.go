// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorTrigger error: %v", err)
		return
	}
	golog.Infof("QueryMonitorTrigger result count: %d", len(result))
}

func TestUpdateMonitorTrigger(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorTrigger found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMonitorTriggerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMonitorTriggerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMonitorTrigger(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorTrigger error: %v", err)
		return
	}
	golog.Infof("UpdateMonitorTrigger result: %s", result.UUID)
}

func TestDeleteMonitorTrigger(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMonitorTrigger is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMonitorTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorTrigger found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMonitorTrigger(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMonitorTrigger error: %v", err)
		return
	}
	golog.Infof("DeleteMonitorTrigger succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMonitorTrigger(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMonitorTrigger is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMonitorTriggerParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMonitorTriggerParamDetail{
	// 		Name: "test-monitortrigger",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMonitorTrigger(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMonitorTrigger error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMonitorTrigger result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMonitorTrigger(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMonitorTrigger error: %v", err)
	// }
}
