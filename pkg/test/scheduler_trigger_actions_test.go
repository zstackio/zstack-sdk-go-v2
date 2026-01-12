// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerTrigger result count: %d", len(result))
}
func TestGetSchedulerTrigger(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestGetSchedulerTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerTrigger found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSchedulerTrigger(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("GetSchedulerTrigger result: %s", result.UUID)
}

func TestUpdateSchedulerTrigger(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerTrigger found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSchedulerTriggerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSchedulerTriggerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSchedulerTrigger(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("UpdateSchedulerTrigger result: %s", result.UUID)
}

func TestDeleteSchedulerTrigger(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSchedulerTrigger is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSchedulerTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerTrigger found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSchedulerTrigger(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("DeleteSchedulerTrigger succeeded for UUID: %s", list[0].UUID)
}

func TestCreateSchedulerTrigger(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSchedulerTrigger is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSchedulerTriggerParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSchedulerTriggerParamDetail{
	// 		Name: "test-schedulertrigger",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSchedulerTrigger(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSchedulerTrigger error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSchedulerTrigger result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSchedulerTrigger(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSchedulerTrigger error: %v", err)
	// }
}
