// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserGroup error: %v", err)
		return
	}
	golog.Infof("QueryUserGroup result count: %d", len(result))
}
func TestGetUserGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetUserGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetUserGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUserGroup error: %v", err)
		return
	}
	golog.Infof("GetUserGroup result: %s", result.UUID)
}

func TestUpdateUserGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateUserGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateUserGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateUserGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateUserGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateUserGroup error: %v", err)
		return
	}
	golog.Infof("UpdateUserGroup result: %s", result.UUID)
}

func TestDeleteUserGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteUserGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteUserGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteUserGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteUserGroup error: %v", err)
		return
	}
	golog.Infof("DeleteUserGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateUserGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateUserGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateUserGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateUserGroupParamDetail{
	// 		Name: "test-usergroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateUserGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateUserGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateUserGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteUserGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteUserGroup error: %v", err)
	// }
}
