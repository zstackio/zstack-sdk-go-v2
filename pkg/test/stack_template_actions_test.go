// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryStackTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryStackTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryStackTemplate error: %v", err)
		return
	}
	golog.Infof("QueryStackTemplate result count: %d", len(result))
}

func TestUpdateStackTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryStackTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateStackTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No StackTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateStackTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateStackTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateStackTemplate(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateStackTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateStackTemplate result: %s", result.Uuid)
}

func TestDeleteStackTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteStackTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryStackTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteStackTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No StackTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteStackTemplate(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteStackTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteStackTemplate succeeded for UUID: %s", list[0].Uuid)
}

func TestAddStackTemplate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddStackTemplate requires valid creation parameters")

}
