// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2ProjectRole(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectRole error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectRole result count: %d", len(result))
}

func TestCreateIAM2ProjectRole(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2ProjectRole is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2ProjectRoleParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2ProjectRoleParamDetail{
	// 		Name: "test-iam2projectrole",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2ProjectRole(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2ProjectRole error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2ProjectRole result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2ProjectRole(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2ProjectRole error: %v", err)
	// }
}
