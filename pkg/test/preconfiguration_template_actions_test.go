// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPreconfigurationTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPreconfigurationTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPreconfigurationTemplate error: %v", err)
		return
	}
	golog.Infof("QueryPreconfigurationTemplate result count: %d", len(result))
}

func TestUpdatePreconfigurationTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPreconfigurationTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePreconfigurationTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PreconfigurationTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePreconfigurationTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePreconfigurationTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePreconfigurationTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePreconfigurationTemplate error: %v", err)
		return
	}
	golog.Infof("UpdatePreconfigurationTemplate result: %s", result.UUID)
}

func TestDeletePreconfigurationTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePreconfigurationTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPreconfigurationTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePreconfigurationTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PreconfigurationTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeletePreconfigurationTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePreconfigurationTemplate error: %v", err)
		return
	}
	golog.Infof("DeletePreconfigurationTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestAddPreconfigurationTemplate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddPreconfigurationTemplate requires valid creation parameters")

}
