// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJobGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerJobGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJobGroup error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJobGroup result count: %d", len(result))
}

func TestUpdateSchedulerJobGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJobGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerJobGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJobGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSchedulerJobGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSchedulerJobGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSchedulerJobGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerJobGroup error: %v", err)
		return
	}
	golog.Infof("UpdateSchedulerJobGroup result: %s", result.UUID)
}

func TestDeleteSchedulerJobGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSchedulerJobGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJobGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSchedulerJobGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJobGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSchedulerJobGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSchedulerJobGroup error: %v", err)
		return
	}
	golog.Infof("DeleteSchedulerJobGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateSchedulerJobGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSchedulerJobGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSchedulerJobGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSchedulerJobGroupParamDetail{
	// 		Name: "test-schedulerjobgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSchedulerJobGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSchedulerJobGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSchedulerJobGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSchedulerJobGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSchedulerJobGroup error: %v", err)
	// }
}
