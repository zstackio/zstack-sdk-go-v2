// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmNic(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmNic error: %v", err)
		return
	}
	golog.Infof("QueryVmNic result count: %d", len(result))
}
func TestGetVmNic(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmNic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmNic found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmNic(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmNic error: %v", err)
		return
	}
	golog.Infof("GetVmNic result: %s", result.UUID)
}

func TestDeleteVmNic(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVmNic is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVmNic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmNic found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVmNic(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVmNic error: %v", err)
		return
	}
	golog.Infof("DeleteVmNic succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVmNic(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVmNic is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVmNicParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVmNicParamDetail{
	// 		Name: "test-vmnic",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVmNic(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVmNic error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVmNic result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVmNic(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVmNic error: %v", err)
	// }
}
