// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VxlanNetwork result count: %d", len(result))
}
func TestGetL2VxlanNetwork(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VxlanNetwork found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL2VxlanNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("GetL2VxlanNetwork result: %s", result.UUID)
}

func TestCreateL2VxlanNetwork(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateL2VxlanNetwork is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateL2VxlanNetworkParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateL2VxlanNetworkParamDetail{
	// 		Name: "test-l2vxlannetwork",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateL2VxlanNetwork(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateL2VxlanNetwork error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateL2VxlanNetwork result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteL2VxlanNetwork(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteL2VxlanNetwork error: %v", err)
	// }
}
