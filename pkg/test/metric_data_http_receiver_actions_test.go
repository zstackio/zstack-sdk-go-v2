// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMetricDataHttpReceiver(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMetricDataHttpReceiver(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMetricDataHttpReceiver error: %v", err)
		return
	}
	golog.Infof("QueryMetricDataHttpReceiver result count: %d", len(result))
}
func TestGetMetricDataHttpReceiver(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMetricDataHttpReceiver(&queryParam)
	if err != nil {
		t.Errorf("TestGetMetricDataHttpReceiver Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MetricDataHttpReceiver found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMetricDataHttpReceiver(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMetricDataHttpReceiver error: %v", err)
		return
	}
	golog.Infof("GetMetricDataHttpReceiver result: %s", result.UUID)
}

func TestDeleteMetricDataHttpReceiver(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMetricDataHttpReceiver is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMetricDataHttpReceiver(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMetricDataHttpReceiver Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MetricDataHttpReceiver found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMetricDataHttpReceiver(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMetricDataHttpReceiver error: %v", err)
		return
	}
	golog.Infof("DeleteMetricDataHttpReceiver succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMetricDataHttpReceiver(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMetricDataHttpReceiver is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMetricDataHttpReceiverParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMetricDataHttpReceiverParamDetail{
	// 		Name: "test-metricdatahttpreceiver",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMetricDataHttpReceiver(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMetricDataHttpReceiver error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMetricDataHttpReceiver result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMetricDataHttpReceiver(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMetricDataHttpReceiver error: %v", err)
	// }
}
