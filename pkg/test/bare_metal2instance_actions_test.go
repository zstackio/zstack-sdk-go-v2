// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Instance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Instance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Instance error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Instance result count: %d", len(result))
}

func TestUpdateBareMetal2Instance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Instance(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Instance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Instance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2InstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2InstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2Instance(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Instance error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2Instance result: %s", result.UUID)
}

func TestCreateBareMetal2Instance(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBareMetal2Instance is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBareMetal2InstanceParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBareMetal2InstanceParamDetail{
	// 		Name: "test-baremetal2instance",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBareMetal2Instance(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBareMetal2Instance error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBareMetal2Instance result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBareMetal2Instance(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBareMetal2Instance error: %v", err)
	// }
}

func TestReconnectBareMetal2Instance(t *testing.T) {
	// ReconnectBareMetal2Instance operation
	t.Skip("TestReconnectBareMetal2Instance requires manual implementation")

}

func TestStartBareMetal2Instance(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartBareMetal2Instance requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QueryBareMetal2Instance(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped BareMetal2Instance found")
	// 	return
	// }
	// startParam := param.StartBareMetal2InstanceParam{}
	// result, err := accountLoginCli.StartBareMetal2Instance(list[0].Uuid, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartBareMetal2Instance error: %v", err)
	// }
	// golog.Infof("StartBareMetal2Instance result: %v", result)

}
