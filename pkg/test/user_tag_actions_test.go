// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUserTag(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserTag error: %v", err)
		return
	}
	golog.Infof("QueryUserTag result count: %d", len(result))
}
func TestGetUserTag(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryUserTag(&queryParam)
	if err != nil {
		t.Errorf("TestGetUserTag Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserTag found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetUserTag(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUserTag error: %v", err)
		return
	}
	golog.Infof("GetUserTag result: %s", result.UUID)
}

func TestCreateUserTag(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateUserTag is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateUserTagParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateUserTagParamDetail{
	// 		Name: "test-usertag",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateUserTag(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateUserTag error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateUserTag result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteUserTag(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteUserTag error: %v", err)
	// }
}
