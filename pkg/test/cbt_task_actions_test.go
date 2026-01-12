// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCbtTask(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCbtTask(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCbtTask error: %v", err)
		return
	}
	golog.Infof("QueryCbtTask result count: %d", len(result))
}

func TestDeleteCbtTask(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCbtTask is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCbtTask(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCbtTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CbtTask found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCbtTask(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCbtTask error: %v", err)
		return
	}
	golog.Infof("DeleteCbtTask succeeded for UUID: %s", list[0].UUID)
}

func TestCreateCbtTask(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateCbtTask is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateCbtTaskParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateCbtTaskParamDetail{
	// 		Name: "test-cbttask",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateCbtTask(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateCbtTask error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateCbtTask result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteCbtTask(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteCbtTask error: %v", err)
	// }
}
