// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVRouterRouteTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("QueryVRouterRouteTable result count: %d", len(result))
}

func TestGetVRouterRouteTable(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestGetVRouterRouteTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VRouterRouteTable found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVRouterRouteTable(list[0].Uuid)
	if err != nil {
		t.Errorf("TestGetVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("GetVRouterRouteTable result: %s", result.Uuid)
}

func TestUpdateVRouterRouteTable(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVRouterRouteTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VRouterRouteTable found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVRouterRouteTableParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVRouterRouteTableParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVRouterRouteTable(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("UpdateVRouterRouteTable result: %s", result.Uuid)
}

func TestDeleteVRouterRouteTable(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVRouterRouteTable is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVRouterRouteTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VRouterRouteTable found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVRouterRouteTable(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("DeleteVRouterRouteTable succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateVRouterRouteTable(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVRouterRouteTable is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVRouterRouteTableParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVRouterRouteTableParamDetail{
	// 		Name: "test-vrouterroutetable",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVRouterRouteTable(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVRouterRouteTable error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVRouterRouteTable result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVRouterRouteTable(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVRouterRouteTable error: %v", err)
	// }
}
