// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmCdRom(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmCdRom(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmCdRom error: %v", err)
		return
	}
	golog.Infof("QueryVmCdRom result count: %d", len(result))
}
func TestGetVmCdRom(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmCdRom(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmCdRom Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmCdRom found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmCdRom(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmCdRom error: %v", err)
		return
	}
	golog.Infof("GetVmCdRom result: %s", result.UUID)
}

func TestUpdateVmCdRom(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmCdRom(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVmCdRom Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmCdRom found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVmCdRomParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVmCdRomParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVmCdRom(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVmCdRom error: %v", err)
		return
	}
	golog.Infof("UpdateVmCdRom result: %s", result.UUID)
}

func TestDeleteVmCdRom(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVmCdRom is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmCdRom(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVmCdRom Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmCdRom found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVmCdRom(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVmCdRom error: %v", err)
		return
	}
	golog.Infof("DeleteVmCdRom succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVmCdRom(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVmCdRom is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVmCdRomParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVmCdRomParamDetail{
	// 		Name: "test-vmcdrom",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVmCdRom(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVmCdRom error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVmCdRom result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVmCdRom(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVmCdRom error: %v", err)
	// }
}
