// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryVmInstance result count: %d", len(result))
}

func TestUpdateVmInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVmInstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVmInstance(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVmInstance error: %v", err)
		return
	}
	golog.Infof("UpdateVmInstance result: %s", result.Uuid)
}

func TestCreateVmInstance(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVmInstance is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVmInstanceParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVmInstanceParamDetail{
	// 		Name: "test-vminstance",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVmInstance(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVmInstance error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVmInstance result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVmInstance(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVmInstance error: %v", err)
	// }
}

func TestCloneVmInstance(t *testing.T) {
	// Clone operation
	t.Skip("TestCloneVmInstance requires a valid resource to clone")

}

func TestResumeVmInstance(t *testing.T) {
	// ResumeVmInstance operation
	t.Skip("TestResumeVmInstance requires manual implementation")

}

func TestStartVmInstance(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartVmInstance requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.Query(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped  found")
	// 	return
	// }
	// startParam := param.StartVmInstanceParam{}
	// result, err := accountLoginCli.StartVmInstance(list[0].Uuid, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartVmInstance error: %v", err)
	// }
	// golog.Infof("StartVmInstance result: %v", result)

}

func TestStopVmInstance(t *testing.T) {
	// Stop operation - requires a running resource UUID
	t.Skip("TestStopVmInstance requires a running resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Running")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.Query(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No running  found")
	// 	return
	// }
	// stopParam := param.StopVmInstanceParam{}
	// result, err := accountLoginCli.StopVmInstance(list[0].Uuid, stopParam)
	// if err != nil {
	// 	t.Errorf("TestStopVmInstance error: %v", err)
	// }
	// golog.Infof("StopVmInstance result: %v", result)

}

func TestExpungeVmInstance(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeVmInstance is dangerous - permanently deletes resource")

}

func TestRebootVmInstance(t *testing.T) {
	// Reboot operation - requires a running resource
	t.Skip("TestRebootVmInstance requires a running resource UUID")

}

func TestDestroyVmInstance(t *testing.T) {
	// DestroyVmInstance operation
	t.Skip("TestDestroyVmInstance requires manual implementation")

}

func TestRecoverVmInstance(t *testing.T) {
	// Recover operation - requires a deleted resource
	t.Skip("TestRecoverVmInstance requires a deleted resource UUID")

}
