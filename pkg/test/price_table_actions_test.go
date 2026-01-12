// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPriceTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPriceTable error: %v", err)
		return
	}
	golog.Infof("QueryPriceTable result count: %d", len(result))
}
func TestGetPriceTable(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestGetPriceTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PriceTable found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPriceTable(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPriceTable error: %v", err)
		return
	}
	golog.Infof("GetPriceTable result: %s", result.UUID)
}

func TestUpdatePriceTable(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePriceTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PriceTable found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePriceTableParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePriceTableParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePriceTable(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePriceTable error: %v", err)
		return
	}
	golog.Infof("UpdatePriceTable result: %s", result.UUID)
}

func TestDeletePriceTable(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePriceTable is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePriceTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PriceTable found to test Delete")
		return
	}

	err = accountLoginCli.DeletePriceTable(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePriceTable error: %v", err)
		return
	}
	golog.Infof("DeletePriceTable succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePriceTable(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePriceTable is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePriceTableParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePriceTableParamDetail{
	// 		Name: "test-pricetable",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePriceTable(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePriceTable error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePriceTable result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePriceTable(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePriceTable error: %v", err)
	// }
}
