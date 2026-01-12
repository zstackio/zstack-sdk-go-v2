// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZone(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZone error: %v", err)
		return
	}
	golog.Infof("QueryZone result count: %d", len(result))
}

func TestGetZone(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestGetZone Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Zone found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetZone(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetZone error: %v", err)
		return
	}
	golog.Infof("GetZone result: %s", result.UUID)
}

func TestUpdateZone(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateZone Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Zone found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateZoneParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateZoneParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateZone(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateZone error: %v", err)
		return
	}
	golog.Infof("UpdateZone result: %s", result.UUID)
}

func TestDeleteZone(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteZone is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteZone Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Zone found to test Delete")
		return
	}

	err = accountLoginCli.DeleteZone(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteZone error: %v", err)
		return
	}
	golog.Infof("DeleteZone succeeded for UUID: %s", list[0].UUID)
}

func TestCreateZone(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateZone is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateZoneParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateZoneParamDetail{
	// 		Name: "test-zone",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateZone(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateZone error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateZone result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteZone(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteZone error: %v", err)
	// }
}
