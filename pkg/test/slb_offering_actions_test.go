// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySlbOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySlbOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbOffering error: %v", err)
		return
	}
	golog.Infof("QuerySlbOffering result count: %d", len(result))
}

func TestCreateSlbOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSlbOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSlbOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSlbOfferingParamDetail{
	// 		Name: "test-slboffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSlbOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSlbOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSlbOffering result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSlbOffering(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSlbOffering error: %v", err)
	// }
}
