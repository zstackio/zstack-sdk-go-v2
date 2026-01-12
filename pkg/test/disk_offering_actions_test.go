// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDiskOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDiskOffering error: %v", err)
		return
	}
	golog.Infof("QueryDiskOffering result count: %d", len(result))
}
func TestGetDiskOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetDiskOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DiskOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetDiskOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDiskOffering error: %v", err)
		return
	}
	golog.Infof("GetDiskOffering result: %s", result.UUID)
}

func TestUpdateDiskOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateDiskOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DiskOffering found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateDiskOfferingParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateDiskOfferingParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateDiskOffering(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateDiskOffering error: %v", err)
		return
	}
	golog.Infof("UpdateDiskOffering result: %s", result.UUID)
}

func TestDeleteDiskOffering(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteDiskOffering is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteDiskOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DiskOffering found to test Delete")
		return
	}

	err = accountLoginCli.DeleteDiskOffering(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteDiskOffering error: %v", err)
		return
	}
	golog.Infof("DeleteDiskOffering succeeded for UUID: %s", list[0].UUID)
}

func TestCreateDiskOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateDiskOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateDiskOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateDiskOfferingParamDetail{
	// 		Name: "test-diskoffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateDiskOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateDiskOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateDiskOffering result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteDiskOffering(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteDiskOffering error: %v", err)
	// }
}
