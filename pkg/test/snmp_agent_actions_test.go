// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySnmpAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySnmpAgent(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySnmpAgent error: %v", err)
		return
	}
	golog.Infof("QuerySnmpAgent result count: %d", len(result))
}
func TestGetSnmpAgent(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySnmpAgent(&queryParam)
	if err != nil {
		t.Errorf("TestGetSnmpAgent Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SnmpAgent found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSnmpAgent(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSnmpAgent error: %v", err)
		return
	}
	golog.Infof("GetSnmpAgent result: %s", result.UUID)
}

func TestUpdateSnmpAgent(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySnmpAgent(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSnmpAgent Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SnmpAgent found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSnmpAgentParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSnmpAgentParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSnmpAgent(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSnmpAgent error: %v", err)
		return
	}
	golog.Infof("UpdateSnmpAgent result: %s", result.UUID)
}

func TestCreateSnmpAgent(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSnmpAgent is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSnmpAgentParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSnmpAgentParamDetail{
	// 		Name: "test-snmpagent",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSnmpAgent(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSnmpAgent error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSnmpAgent result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSnmpAgent(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSnmpAgent error: %v", err)
	// }
}

func TestStartSnmpAgent(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartSnmpAgent requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QuerySnmpAgent(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped SnmpAgent found")
	// 	return
	// }
	// startParam := param.StartSnmpAgentParam{}
	// result, err := accountLoginCli.StartSnmpAgent(list[0].UUID, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartSnmpAgent error: %v", err)
	// }
	// golog.Infof("StartSnmpAgent result: %v", result)

}

func TestStopSnmpAgent(t *testing.T) {
	// Stop operation - requires a running resource UUID
	t.Skip("TestStopSnmpAgent requires a running resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Running")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QuerySnmpAgent(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No running SnmpAgent found")
	// 	return
	// }
	// stopParam := param.StopSnmpAgentParam{}
	// result, err := accountLoginCli.StopSnmpAgent(list[0].UUID, stopParam)
	// if err != nil {
	// 	t.Errorf("TestStopSnmpAgent error: %v", err)
	// }
	// golog.Infof("StopSnmpAgent result: %v", result)

}
