// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryRole(&queryParam)
	if err != nil {
		t.Errorf("TestQueryRole error: %v", err)
		return
	}
	golog.Infof("QueryRole result count: %d", len(result))
}

func TestUpdateRole(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryRole(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateRole Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Role found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateRoleParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateRoleParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateRole(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateRole error: %v", err)
		return
	}
	golog.Infof("UpdateRole result: %s", result.UUID)
}

func TestDeleteRole(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteRole is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryRole(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteRole Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Role found to test Delete")
		return
	}

	err = accountLoginCli.DeleteRole(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteRole error: %v", err)
		return
	}
	golog.Infof("DeleteRole succeeded for UUID: %s", list[0].UUID)
}

func TestCreateRole(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateRole is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateRoleParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateRoleParamDetail{
	// 		Name: "test-role",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateRole(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateRole error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateRole result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteRole(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteRole error: %v", err)
	// }
}
