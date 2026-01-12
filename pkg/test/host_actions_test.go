// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHost error: %v", err)
		return
	}
	golog.Infof("QueryHost result count: %d", len(result))
}

func TestUpdateHost(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Host found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateHostParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateHostParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateHost(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateHost error: %v", err)
		return
	}
	golog.Infof("UpdateHost result: %s", result.UUID)
}

func TestDeleteHost(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHost is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Host found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHost(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHost error: %v", err)
		return
	}
	golog.Infof("DeleteHost succeeded for UUID: %s", list[0].UUID)
}

func TestReconnectHost(t *testing.T) {
	// ReconnectHost operation
	t.Skip("TestReconnectHost requires manual implementation")

}
