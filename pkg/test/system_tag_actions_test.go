// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySystemTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySystemTag error: %v", err)
		return
	}
	golog.Infof("QuerySystemTag result count: %d", len(result))
}

func TestUpdateSystemTag(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSystemTag Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SystemTag found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSystemTagParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSystemTagParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSystemTag(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSystemTag error: %v", err)
		return
	}
	golog.Infof("UpdateSystemTag result: %s", result.Uuid)
}

func TestCreateSystemTag(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSystemTag is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSystemTagParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSystemTagParamDetail{
	// 		Name: "test-systemtag",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSystemTag(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSystemTag error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSystemTag result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSystemTag(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSystemTag error: %v", err)
	// }
}
