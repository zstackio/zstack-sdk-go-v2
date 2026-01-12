// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VlanNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VlanNetwork result count: %d", len(result))
}

func TestCreateL2VlanNetwork(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateL2VlanNetwork is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateL2VlanNetworkParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateL2VlanNetworkParamDetail{
	// 		Name: "test-l2vlannetwork",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateL2VlanNetwork(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateL2VlanNetwork error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateL2VlanNetwork result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteL2VlanNetwork(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteL2VlanNetwork error: %v", err)
	// }
}
