// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryManagementNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryManagementNode error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryManagementNode result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.HostName)
	}
	golog.Infof("======================================")
}

func TestPageManagementNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestPageManagementNode error: %v", err)
		return
	}
	golog.Infof("PageManagementNode result: total=%d, returned=%d", total, len(result))
}

func TestGetManagementNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestGetManagementNode Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ManagementNode found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetManagementNode(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetManagementNode error: %v", err)
		return
	}
	golog.Infof("GetManagementNode result: %s, HostName: %s", result.UUID, result.HostName)
}
