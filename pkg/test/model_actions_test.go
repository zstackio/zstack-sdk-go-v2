// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModel(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModel(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModel error: %v", err)
		return
	}
	golog.Infof("QueryModel result count: %d", len(result))
}

func TestUpdateModel(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModel(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateModel Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Model found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateModelParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateModelParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateModel(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateModel error: %v", err)
		return
	}
	golog.Infof("UpdateModel result: %s", result.UUID)
}

func TestDeleteModel(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteModel is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModel(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteModel Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Model found to test Delete")
		return
	}

	err = accountLoginCli.DeleteModel(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteModel error: %v", err)
		return
	}
	golog.Infof("DeleteModel succeeded for UUID: %s", list[0].UUID)
}

func TestAddModel(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddModel requires valid creation parameters")

}
