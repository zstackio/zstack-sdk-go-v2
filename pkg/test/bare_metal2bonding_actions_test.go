// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Bonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Bonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Bonding error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Bonding result count: %d", len(result))
}

func TestCreateBareMetal2Bonding(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBareMetal2Bonding is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBareMetal2BondingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBareMetal2BondingParamDetail{
	// 		Name: "test-baremetal2bonding",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBareMetal2Bonding(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBareMetal2Bonding error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBareMetal2Bonding result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBareMetal2Bonding(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBareMetal2Bonding error: %v", err)
	// }
}
