// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBaremetalInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalInstance error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalInstance result count: %d", len(result))
}

func TestUpdateBaremetalInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBaremetalInstance(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BaremetalInstance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBaremetalInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBaremetalInstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBaremetalInstance(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalInstance error: %v", err)
		return
	}
	golog.Infof("UpdateBaremetalInstance result: %s", result.UUID)
}

func TestCreateBaremetalInstance(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBaremetalInstance is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBaremetalInstanceParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBaremetalInstanceParamDetail{
	// 		Name: "test-baremetalinstance",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBaremetalInstance(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBaremetalInstance error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBaremetalInstance result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBaremetalInstance(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBaremetalInstance error: %v", err)
	// }
}

func TestRebootBaremetalInstance(t *testing.T) {
	// Reboot operation - requires a running resource
	t.Skip("TestRebootBaremetalInstance requires a running resource UUID")

}

func TestStartBaremetalInstance(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartBaremetalInstance requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.Query(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped  found")
	// 	return
	// }
	// startParam := param.StartBaremetalInstanceParam{}
	// result, err := accountLoginCli.StartBaremetalInstance(list[0].Uuid, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartBaremetalInstance error: %v", err)
	// }
	// golog.Infof("StartBaremetalInstance result: %v", result)

}

func TestDestroyBaremetalInstance(t *testing.T) {
	// DestroyBaremetalInstance operation
	t.Skip("TestDestroyBaremetalInstance requires manual implementation")

}

func TestExpungeBaremetalInstance(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeBaremetalInstance is dangerous - permanently deletes resource")

}

func TestStopBaremetalInstance(t *testing.T) {
	// Stop operation - requires a running resource UUID
	t.Skip("TestStopBaremetalInstance requires a running resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Running")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QueryBaremetalInstance(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No running BaremetalInstance found")
	// 	return
	// }
	// stopParam := param.StopBaremetalInstanceParam{}
	// result, err := accountLoginCli.StopBaremetalInstance(list[0].Uuid, stopParam)
	// if err != nil {
	// 	t.Errorf("TestStopBaremetalInstance error: %v", err)
	// }
	// golog.Infof("StopBaremetalInstance result: %v", result)

}

func TestRecoverBaremetalInstance(t *testing.T) {
	// Recover operation - requires a deleted resource
	t.Skip("TestRecoverBaremetalInstance requires a deleted resource UUID")

}
