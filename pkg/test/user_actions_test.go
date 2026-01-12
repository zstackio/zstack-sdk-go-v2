// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUser(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUser error: %v", err)
		return
	}
	golog.Infof("QueryUser result count: %d", len(result))
}

func TestUpdateUser(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateUser Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No User found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateUserParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateUserParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateUser(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateUser error: %v", err)
		return
	}
	golog.Infof("UpdateUser result: %s", result.Uuid)
}

func TestDeleteUser(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteUser is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteUser Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No User found to test Delete")
		return
	}

	err = accountLoginCli.DeleteUser(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteUser error: %v", err)
		return
	}
	golog.Infof("DeleteUser succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateUser(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateUser is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateUserParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateUserParamDetail{
	// 		Name: "test-user",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateUser(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateUser error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateUser result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteUser(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteUser error: %v", err)
	// }
}
