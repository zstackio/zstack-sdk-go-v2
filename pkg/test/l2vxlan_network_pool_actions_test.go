// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetworkPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("QueryL2VxlanNetworkPool result count: %d", len(result))
}
func TestGetL2VxlanNetworkPool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetworkPool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VxlanNetworkPool found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL2VxlanNetworkPool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("GetL2VxlanNetworkPool result: %s", result.UUID)
}

func TestCreateL2VxlanNetworkPool(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateL2VxlanNetworkPool is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateL2VxlanNetworkPoolParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateL2VxlanNetworkPoolParamDetail{
	// 		Name: "test-l2vxlannetworkpool",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateL2VxlanNetworkPool(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateL2VxlanNetworkPool error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateL2VxlanNetworkPool result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteL2VxlanNetworkPool(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteL2VxlanNetworkPool error: %v", err)
	// }
}
