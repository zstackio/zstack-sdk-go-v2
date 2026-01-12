// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessKey error: %v", err)
		return
	}
	golog.Infof("QueryAccessKey result count: %d", len(result))
}
func TestGetAccessKey(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAccessKey(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccessKey error: %v", err)
		return
	}
	golog.Infof("GetAccessKey result: %s", result.UUID)
}

func TestDeleteAccessKey(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAccessKey is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAccessKey(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAccessKey error: %v", err)
		return
	}
	golog.Infof("DeleteAccessKey succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAccessKey(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAccessKey is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAccessKeyParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAccessKeyParamDetail{
	// 		Name: "test-accesskey",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAccessKey(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAccessKey error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAccessKey result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAccessKey(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAccessKey error: %v", err)
	// }
}
