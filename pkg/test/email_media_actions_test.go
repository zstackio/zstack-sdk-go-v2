// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEmailMedia(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEmailMedia(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEmailMedia error: %v", err)
		return
	}
	golog.Infof("QueryEmailMedia result count: %d", len(result))
}

func TestUpdateEmailMedia(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEmailMedia(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateEmailMedia Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EmailMedia found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEmailMediaParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEmailMediaParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEmailMedia(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEmailMedia error: %v", err)
		return
	}
	golog.Infof("UpdateEmailMedia result: %s", result.UUID)
}

func TestCreateEmailMedia(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateEmailMedia is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateEmailMediaParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateEmailMediaParamDetail{
	// 		Name: "test-emailmedia",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateEmailMedia(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateEmailMedia error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateEmailMedia result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteEmailMedia(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteEmailMedia error: %v", err)
	// }
}
