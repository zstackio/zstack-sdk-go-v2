// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDataset(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDataset(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDataset error: %v", err)
		return
	}
	golog.Infof("QueryDataset result count: %d", len(result))
}
func TestGetDataset(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDataset(&queryParam)
	if err != nil {
		t.Errorf("TestGetDataset Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Dataset found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetDataset(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDataset error: %v", err)
		return
	}
	golog.Infof("GetDataset result: %s", result.UUID)
}

func TestUpdateDataset(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDataset(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateDataset Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Dataset found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateDatasetParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateDatasetParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateDataset(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateDataset error: %v", err)
		return
	}
	golog.Infof("UpdateDataset result: %s", result.UUID)
}

func TestDeleteDataset(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteDataset is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDataset(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteDataset Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Dataset found to test Delete")
		return
	}

	err = accountLoginCli.DeleteDataset(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteDataset error: %v", err)
		return
	}
	golog.Infof("DeleteDataset succeeded for UUID: %s", list[0].UUID)
}

func TestCreateDataset(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateDataset is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateDatasetParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateDatasetParamDetail{
	// 		Name: "test-dataset",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateDataset(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateDataset error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateDataset result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteDataset(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteDataset error: %v", err)
	// }
}
