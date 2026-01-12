// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcHaGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcHaGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcHaGroup error: %v", err)
		return
	}
	golog.Infof("QueryVpcHaGroup result count: %d", len(result))
}
func TestGetVpcHaGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcHaGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVpcHaGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcHaGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVpcHaGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVpcHaGroup error: %v", err)
		return
	}
	golog.Infof("GetVpcHaGroup result: %s", result.UUID)
}

func TestUpdateVpcHaGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcHaGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVpcHaGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcHaGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVpcHaGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVpcHaGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVpcHaGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVpcHaGroup error: %v", err)
		return
	}
	golog.Infof("UpdateVpcHaGroup result: %s", result.UUID)
}

func TestDeleteVpcHaGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVpcHaGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcHaGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVpcHaGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcHaGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVpcHaGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVpcHaGroup error: %v", err)
		return
	}
	golog.Infof("DeleteVpcHaGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVpcHaGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVpcHaGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVpcHaGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVpcHaGroupParamDetail{
	// 		Name: "test-vpchagroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVpcHaGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVpcHaGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVpcHaGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVpcHaGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVpcHaGroup error: %v", err)
	// }
}
