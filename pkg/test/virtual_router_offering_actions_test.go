// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterOffering result count: %d", len(result))
}

func TestUpdateVirtualRouterOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVirtualRouterOffering(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVirtualRouterOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VirtualRouterOffering found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVirtualRouterOfferingParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVirtualRouterOfferingParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVirtualRouterOffering(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("UpdateVirtualRouterOffering result: %s", result.Uuid)
}

func TestCreateVirtualRouterOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVirtualRouterOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVirtualRouterOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVirtualRouterOfferingParamDetail{
	// 		Name: "test-virtualrouteroffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVirtualRouterOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVirtualRouterOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVirtualRouterOffering result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVirtualRouterOffering(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVirtualRouterOffering error: %v", err)
	// }
}
