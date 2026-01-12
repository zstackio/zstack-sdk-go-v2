// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryConsoleProxyAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestQueryConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("QueryConsoleProxyAgent result count: %d", len(result))
}
func TestGetConsoleProxyAgent(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestGetConsoleProxyAgent Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ConsoleProxyAgent found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetConsoleProxyAgent(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("GetConsoleProxyAgent result: %s", result.UUID)
}

func TestUpdateConsoleProxyAgent(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateConsoleProxyAgent Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ConsoleProxyAgent found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateConsoleProxyAgentParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateConsoleProxyAgentParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateConsoleProxyAgent(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("UpdateConsoleProxyAgent result: %s", result.UUID)
}

func TestReconnectConsoleProxyAgent(t *testing.T) {
	// ReconnectConsoleProxyAgent operation
	t.Skip("TestReconnectConsoleProxyAgent requires manual implementation")

}
