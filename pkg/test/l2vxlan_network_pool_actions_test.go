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
