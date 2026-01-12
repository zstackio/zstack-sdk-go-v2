// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserProxyConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUserProxyConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserProxyConfig error: %v", err)
		return
	}
	golog.Infof("QueryUserProxyConfig result count: %d", len(result))
}
func TestGetUserProxyConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserProxyConfig(&queryParam)
	if err != nil {
		t.Errorf("TestGetUserProxyConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserProxyConfig found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetUserProxyConfig(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUserProxyConfig error: %v", err)
		return
	}
	golog.Infof("GetUserProxyConfig result: %s", result.UUID)
}

func TestUpdateUserProxyConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserProxyConfig(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateUserProxyConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserProxyConfig found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateUserProxyConfigParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateUserProxyConfigParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateUserProxyConfig(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateUserProxyConfig error: %v", err)
		return
	}
	golog.Infof("UpdateUserProxyConfig result: %s", result.UUID)
}

func TestDeleteUserProxyConfig(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteUserProxyConfig is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserProxyConfig(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteUserProxyConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserProxyConfig found to test Delete")
		return
	}

	err = accountLoginCli.DeleteUserProxyConfig(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteUserProxyConfig error: %v", err)
		return
	}
	golog.Infof("DeleteUserProxyConfig succeeded for UUID: %s", list[0].UUID)
}

func TestCreateUserProxyConfig(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateUserProxyConfig is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateUserProxyConfigParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateUserProxyConfigParamDetail{
	// 		Name: "test-userproxyconfig",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateUserProxyConfig(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateUserProxyConfig error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateUserProxyConfig result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteUserProxyConfig(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteUserProxyConfig error: %v", err)
	// }
}
