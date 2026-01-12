// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPciDeviceOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPciDeviceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDeviceOffering error: %v", err)
		return
	}
	golog.Infof("QueryPciDeviceOffering result count: %d", len(result))
}
func TestGetPciDeviceOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPciDeviceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetPciDeviceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDeviceOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPciDeviceOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPciDeviceOffering error: %v", err)
		return
	}
	golog.Infof("GetPciDeviceOffering result: %s", result.UUID)
}

func TestDeletePciDeviceOffering(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePciDeviceOffering is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPciDeviceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePciDeviceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDeviceOffering found to test Delete")
		return
	}

	err = accountLoginCli.DeletePciDeviceOffering(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePciDeviceOffering error: %v", err)
		return
	}
	golog.Infof("DeletePciDeviceOffering succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePciDeviceOffering(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePciDeviceOffering is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePciDeviceOfferingParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePciDeviceOfferingParamDetail{
	// 		Name: "test-pcideviceoffering",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePciDeviceOffering(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePciDeviceOffering error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePciDeviceOffering result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePciDeviceOffering(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePciDeviceOffering error: %v", err)
	// }
}
