// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryManagementNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryManagementNode error: %v", err)
		return
	}
	golog.Infof("QueryManagementNode result count: %d", len(result))
}
func TestGetManagementNode(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestGetManagementNode Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ManagementNode found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetManagementNode(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetManagementNode error: %v", err)
		return
	}
	golog.Infof("GetManagementNode result: %s", result.UUID)
}
