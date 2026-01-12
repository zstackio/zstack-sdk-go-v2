// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLogServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLogServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLogServer error: %v", err)
		return
	}
	golog.Infof("QueryLogServer result count: %d", len(result))
}

func TestUpdateLogServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLogServer(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLogServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LogServer found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLogServerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLogServerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLogServer(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLogServer error: %v", err)
		return
	}
	golog.Infof("UpdateLogServer result: %s", result.UUID)
}

func TestDeleteLogServer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLogServer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLogServer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLogServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LogServer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLogServer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLogServer error: %v", err)
		return
	}
	golog.Infof("DeleteLogServer succeeded for UUID: %s", list[0].UUID)
}

func TestAddLogServer(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddLogServer requires valid creation parameters")

}
