// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerJob(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJob error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJob result count: %d", len(result))
}
func TestGetSchedulerJob(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJob(&queryParam)
	if err != nil {
		t.Errorf("TestGetSchedulerJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJob found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSchedulerJob(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSchedulerJob error: %v", err)
		return
	}
	golog.Infof("GetSchedulerJob result: %s", result.UUID)
}

func TestUpdateSchedulerJob(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJob(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJob found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSchedulerJobParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSchedulerJobParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSchedulerJob(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSchedulerJob error: %v", err)
		return
	}
	golog.Infof("UpdateSchedulerJob result: %s", result.UUID)
}

func TestDeleteSchedulerJob(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSchedulerJob is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJob(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSchedulerJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJob found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSchedulerJob(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSchedulerJob error: %v", err)
		return
	}
	golog.Infof("DeleteSchedulerJob succeeded for UUID: %s", list[0].UUID)
}

func TestCreateSchedulerJob(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSchedulerJob is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSchedulerJobParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSchedulerJobParamDetail{
	// 		Name: "test-schedulerjob",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSchedulerJob(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSchedulerJob error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSchedulerJob result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSchedulerJob(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSchedulerJob error: %v", err)
	// }
}
