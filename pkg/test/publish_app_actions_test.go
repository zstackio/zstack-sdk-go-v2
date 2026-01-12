// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPublishApp(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPublishApp(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPublishApp error: %v", err)
		return
	}
	golog.Infof("QueryPublishApp result count: %d", len(result))
}

func TestUpdatePublishApp(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPublishApp(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePublishApp Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PublishApp found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePublishAppParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePublishAppParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePublishApp(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePublishApp error: %v", err)
		return
	}
	golog.Infof("UpdatePublishApp result: %s", result.UUID)
}

func TestDeletePublishApp(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePublishApp is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPublishApp(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePublishApp Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PublishApp found to test Delete")
		return
	}

	err = accountLoginCli.DeletePublishApp(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePublishApp error: %v", err)
		return
	}
	golog.Infof("DeletePublishApp succeeded for UUID: %s", list[0].UUID)
}

func TestPublishApp(t *testing.T) {
	// PublishApp operation
	t.Skip("TestPublishApp requires manual implementation")

}
