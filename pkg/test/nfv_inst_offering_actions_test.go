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
func TestGetNfvInstOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNfvInstOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetNfvInstOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NfvInstOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNfvInstOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNfvInstOffering error: %v", err)
		return
	}
	golog.Infof("GetNfvInstOffering result: %s", result.UUID)
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
