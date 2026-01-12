// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCdpTask(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCdpTask(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCdpTask error: %v", err)
		return
	}
	golog.Infof("QueryCdpTask result count: %d", len(result))
}
func TestGetCdpTask(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpTask(&queryParam)
	if err != nil {
		t.Errorf("TestGetCdpTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpTask found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCdpTask(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCdpTask error: %v", err)
		return
	}
	golog.Infof("GetCdpTask result: %s", result.UUID)
}

func TestUpdateCdpTask(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpTask(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCdpTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpTask found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCdpTaskParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCdpTaskParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCdpTask(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCdpTask error: %v", err)
		return
	}
	golog.Infof("UpdateCdpTask result: %s", result.UUID)
}

func TestDeleteCdpTask(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCdpTask is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpTask(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCdpTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpTask found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCdpTask(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCdpTask error: %v", err)
		return
	}
	golog.Infof("DeleteCdpTask succeeded for UUID: %s", list[0].UUID)
}

func TestCreateCdpTask(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateCdpTask is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateCdpTaskParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateCdpTaskParamDetail{
	// 		Name: "test-cdptask",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateCdpTask(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateCdpTask error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateCdpTask result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteCdpTask(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteCdpTask error: %v", err)
	// }
}
