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
func TestGetSNSTopic(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSTopic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTopic found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSTopic(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSTopic error: %v", err)
		return
	}
	golog.Infof("GetSNSTopic result: %s", result.UUID)
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
	result, err := accountLoginCli.UpdateSNSTopic(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSTopic error: %v", err)
		return
	}
	golog.Infof("UpdateSNSTopic result: %s", result.UUID)
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

	err = accountLoginCli.DeleteSNSTopic(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSNSTopic error: %v", err)
		return
	}
	golog.Infof("DeleteSNSTopic succeeded for UUID: %s", list[0].UUID)
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
	// golog.Infof("CreateSNSTopic result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSTopic(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSTopic error: %v", err)
	// }
}
