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
func TestGetSlbOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySlbOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetSlbOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SlbOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSlbOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSlbOffering error: %v", err)
		return
	}
	golog.Infof("GetSlbOffering result: %s", result.UUID)
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
