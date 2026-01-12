// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNfvInst(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNfvInst(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNfvInst error: %v", err)
		return
	}
	golog.Infof("QueryNfvInst result count: %d", len(result))
}

func TestCreateNfvInst(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateNfvInst is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateNfvInstParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateNfvInstParamDetail{
	// 		Name: "test-nfvinst",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateNfvInst(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateNfvInst error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateNfvInst result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteNfvInst(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteNfvInst error: %v", err)
	// }
}

func TestReconnectNfvInst(t *testing.T) {
	// ReconnectNfvInst operation
	t.Skip("TestReconnectNfvInst requires manual implementation")

}
