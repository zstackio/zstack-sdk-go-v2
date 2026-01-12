// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAgentVersion(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAgentVersion(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAgentVersion error: %v", err)
		return
	}
	golog.Infof("QueryAgentVersion result count: %d", len(result))
}
func TestGetAgentVersion(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAgentVersion(&queryParam)
	if err != nil {
		t.Errorf("TestGetAgentVersion Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AgentVersion found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAgentVersion(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAgentVersion error: %v", err)
		return
	}
	golog.Infof("GetAgentVersion result: %s", result.UUID)
}
