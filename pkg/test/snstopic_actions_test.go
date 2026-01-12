// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSTopic(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTopic error: %v", err)
		return
	}
	golog.Infof("QuerySNSTopic result count: %d", len(result))
}

func TestUpdateSNSTopic(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSTopic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTopic found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSTopicParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSTopicParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSTopic(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSTopic error: %v", err)
		return
	}
	golog.Infof("UpdateSNSTopic result: %s", result.Uuid)
}

func TestDeleteSNSTopic(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSNSTopic is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSNSTopic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTopic found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSNSTopic(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSNSTopic error: %v", err)
		return
	}
	golog.Infof("DeleteSNSTopic succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateSNSTopic(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSTopic is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSTopicParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSTopicParamDetail{
	// 		Name: "test-snstopic",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSTopic(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSTopic error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSTopic result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSTopic(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSTopic error: %v", err)
	// }
}
