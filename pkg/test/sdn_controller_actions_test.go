// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySdnController(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySdnController(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySdnController error: %v", err)
		return
	}
	golog.Infof("QuerySdnController result count: %d", len(result))
}

func TestUpdateSdnController(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySdnController(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSdnController Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SdnController found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSdnControllerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSdnControllerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSdnController(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSdnController error: %v", err)
		return
	}
	golog.Infof("UpdateSdnController result: %s", result.Uuid)
}

func TestRemoveSdnController(t *testing.T) {
	// RemoveSdnController operation
	t.Skip("TestRemoveSdnController requires manual implementation")

}

func TestAddSdnController(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSdnController requires valid creation parameters")

}

func TestChangeSdnController(t *testing.T) {
	// Change operation
	t.Skip("TestChangeSdnController requires specific parameters")

}

func TestReconnectSdnController(t *testing.T) {
	// ReconnectSdnController operation
	t.Skip("TestReconnectSdnController requires manual implementation")

}
