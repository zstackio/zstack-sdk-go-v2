// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccount(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccount error: %v", err)
		return
	}
	golog.Infof("QueryAccount result count: %d", len(result))
}

func TestUpdateAccount(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAccount Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Account found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAccountParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAccountParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAccount(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAccount error: %v", err)
		return
	}
	golog.Infof("UpdateAccount result: %s", result.UUID)
}

func TestDeleteAccount(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAccount is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAccount Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Account found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAccount(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAccount error: %v", err)
		return
	}
	golog.Infof("DeleteAccount succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAccount(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAccount is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAccountParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAccountParamDetail{
	// 		Name: "test-account",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAccount(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAccount error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAccount result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAccount(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAccount error: %v", err)
	// }
}
