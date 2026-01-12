// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcSharedQos(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcSharedQos(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcSharedQos error: %v", err)
		return
	}
	golog.Infof("QueryVpcSharedQos result count: %d", len(result))
}

func TestUpdateVpcSharedQos(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcSharedQos(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVpcSharedQos Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcSharedQos found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVpcSharedQosParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVpcSharedQosParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVpcSharedQos(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVpcSharedQos error: %v", err)
		return
	}
	golog.Infof("UpdateVpcSharedQos result: %s", result.Uuid)
}

func TestDeleteVpcSharedQos(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVpcSharedQos is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcSharedQos(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVpcSharedQos Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcSharedQos found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVpcSharedQos(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVpcSharedQos error: %v", err)
		return
	}
	golog.Infof("DeleteVpcSharedQos succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateVpcSharedQos(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVpcSharedQos is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVpcSharedQosParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVpcSharedQosParamDetail{
	// 		Name: "test-vpcsharedqos",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVpcSharedQos(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVpcSharedQos error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVpcSharedQos result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVpcSharedQos(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVpcSharedQos error: %v", err)
	// }
}
