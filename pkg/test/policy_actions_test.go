// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicy error: %v", err)
		return
	}
	golog.Infof("QueryPolicy result count: %d", len(result))
}

func TestDeletePolicy(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePolicy is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Policy found to test Delete")
		return
	}

	err = accountLoginCli.DeletePolicy(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePolicy error: %v", err)
		return
	}
	golog.Infof("DeletePolicy succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePolicy(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePolicy is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePolicyParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePolicyParamDetail{
	// 		Name: "test-policy",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePolicy(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePolicy error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePolicy result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePolicy(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePolicy error: %v", err)
	// }
}
