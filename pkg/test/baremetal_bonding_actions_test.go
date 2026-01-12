// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBaremetalBonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalBonding error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalBonding result count: %d", len(result))
}

func TestCreateBaremetalBonding(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBaremetalBonding is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBaremetalBondingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBaremetalBondingParamDetail{
	// 		Name: "test-baremetalbonding",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBaremetalBonding(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBaremetalBonding error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBaremetalBonding result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBaremetalBonding(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBaremetalBonding error: %v", err)
	// }
}
