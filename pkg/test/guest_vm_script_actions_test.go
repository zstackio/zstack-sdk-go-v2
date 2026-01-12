// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScript(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScript(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScript error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScript result count: %d", len(result))
}

func TestUpdateGuestVmScript(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGuestVmScript(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateGuestVmScript Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GuestVmScript found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateGuestVmScriptParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateGuestVmScriptParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateGuestVmScript(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateGuestVmScript error: %v", err)
		return
	}
	golog.Infof("UpdateGuestVmScript result: %s", result.UUID)
}

func TestDeleteGuestVmScript(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteGuestVmScript is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGuestVmScript(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteGuestVmScript Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GuestVmScript found to test Delete")
		return
	}

	err = accountLoginCli.DeleteGuestVmScript(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteGuestVmScript error: %v", err)
		return
	}
	golog.Infof("DeleteGuestVmScript succeeded for UUID: %s", list[0].UUID)
}

func TestCreateGuestVmScript(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateGuestVmScript is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateGuestVmScriptParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateGuestVmScriptParamDetail{
	// 		Name: "test-guestvmscript",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateGuestVmScript(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateGuestVmScript error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateGuestVmScript result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteGuestVmScript(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteGuestVmScript error: %v", err)
	// }
}
