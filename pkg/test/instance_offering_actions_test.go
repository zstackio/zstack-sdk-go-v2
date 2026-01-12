// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryInstanceOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryInstanceOffering error: %v", err)
		return
	}
	golog.Infof("QueryInstanceOffering result count: %d", len(result))
}
func TestGetInstanceOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetInstanceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No InstanceOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetInstanceOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetInstanceOffering error: %v", err)
		return
	}
	golog.Infof("GetInstanceOffering result: %s", result.UUID)
}

func TestUpdateInstanceOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateInstanceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No InstanceOffering found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateInstanceOfferingParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateInstanceOfferingParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateInstanceOffering(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateInstanceOffering error: %v", err)
		return
	}
	golog.Infof("UpdateInstanceOffering result: %s", result.UUID)
}

func TestDeleteInstanceOffering(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteInstanceOffering is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteInstanceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No InstanceOffering found to test Delete")
		return
	}

	err = accountLoginCli.DeleteInstanceOffering(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteInstanceOffering error: %v", err)
		return
	}
	golog.Infof("DeleteInstanceOffering succeeded for UUID: %s", list[0].UUID)
}

func TestCreateInstanceOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateInstanceOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateInstanceOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateInstanceOfferingParamDetail{
	// 		Name: "test-instanceoffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateInstanceOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateInstanceOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateInstanceOffering result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteInstanceOffering(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteInstanceOffering error: %v", err)
	// }
}

func TestChangeInstanceOffering(t *testing.T) {
	// Change operation
	t.Skip("TestChangeInstanceOffering requires specific parameters")

}
