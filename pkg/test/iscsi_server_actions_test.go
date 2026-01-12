// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIscsiServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIscsiServer error: %v", err)
		return
	}
	golog.Infof("QueryIscsiServer result count: %d", len(result))
}
func TestGetIscsiServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestGetIscsiServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IscsiServer found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIscsiServer(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIscsiServer error: %v", err)
		return
	}
	golog.Infof("GetIscsiServer result: %s", result.UUID)
}

func TestUpdateIscsiServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIscsiServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IscsiServer found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIscsiServerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIscsiServerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIscsiServer(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIscsiServer error: %v", err)
		return
	}
	golog.Infof("UpdateIscsiServer result: %s", result.UUID)
}

func TestDeleteIscsiServer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIscsiServer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIscsiServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IscsiServer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIscsiServer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIscsiServer error: %v", err)
		return
	}
	golog.Infof("DeleteIscsiServer succeeded for UUID: %s", list[0].UUID)
}

func TestAddIscsiServer(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddIscsiServer requires valid creation parameters")

}

func TestRefreshIscsiServer(t *testing.T) {
	// RefreshIscsiServer operation
	t.Skip("TestRefreshIscsiServer requires manual implementation")

}
