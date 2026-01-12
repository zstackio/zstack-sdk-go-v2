// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNfvInstOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNfvInstOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNfvInstOffering error: %v", err)
		return
	}
	golog.Infof("QueryNfvInstOffering result count: %d", len(result))
}

func TestCreateNfvInstOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateNfvInstOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateNfvInstOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateNfvInstOfferingParamDetail{
	// 		Name: "test-nfvinstoffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateNfvInstOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateNfvInstOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateNfvInstOffering result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteNfvInstOffering(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteNfvInstOffering error: %v", err)
	// }
}
