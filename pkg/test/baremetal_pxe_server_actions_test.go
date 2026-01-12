// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalPxeServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBaremetalPxeServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalPxeServer error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalPxeServer result count: %d", len(result))
}

func TestUpdateBaremetalPxeServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBaremetalPxeServer(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalPxeServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BaremetalPxeServer found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBaremetalPxeServerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBaremetalPxeServerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBaremetalPxeServer(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalPxeServer error: %v", err)
		return
	}
	golog.Infof("UpdateBaremetalPxeServer result: %s", result.UUID)
}

func TestDeleteBaremetalPxeServer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBaremetalPxeServer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBaremetalPxeServer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBaremetalPxeServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BaremetalPxeServer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBaremetalPxeServer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBaremetalPxeServer error: %v", err)
		return
	}
	golog.Infof("DeleteBaremetalPxeServer succeeded for UUID: %s", list[0].UUID)
}

func TestCreateBaremetalPxeServer(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBaremetalPxeServer is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBaremetalPxeServerParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBaremetalPxeServerParamDetail{
	// 		Name: "test-baremetalpxeserver",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBaremetalPxeServer(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBaremetalPxeServer error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBaremetalPxeServer result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBaremetalPxeServer(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBaremetalPxeServer error: %v", err)
	// }
}

func TestStartBaremetalPxeServer(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartBaremetalPxeServer requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QueryBaremetalPxeServer(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped BaremetalPxeServer found")
	// 	return
	// }
	// startParam := param.StartBaremetalPxeServerParam{}
	// result, err := accountLoginCli.StartBaremetalPxeServer(list[0].Uuid, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartBaremetalPxeServer error: %v", err)
	// }
	// golog.Infof("StartBaremetalPxeServer result: %v", result)

}

func TestReconnectBaremetalPxeServer(t *testing.T) {
	// ReconnectBaremetalPxeServer operation
	t.Skip("TestReconnectBaremetalPxeServer requires manual implementation")

}

func TestStopBaremetalPxeServer(t *testing.T) {
	// Stop operation - requires a running resource UUID
	t.Skip("TestStopBaremetalPxeServer requires a running resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Running")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QueryBaremetalPxeServer(&queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No running BaremetalPxeServer found")
	// 	return
	// }
	// stopParam := param.StopBaremetalPxeServerParam{}
	// result, err := accountLoginCli.StopBaremetalPxeServer(list[0].Uuid, stopParam)
	// if err != nil {
	// 	t.Errorf("TestStopBaremetalPxeServer error: %v", err)
	// }
	// golog.Infof("StopBaremetalPxeServer result: %v", result)

}
