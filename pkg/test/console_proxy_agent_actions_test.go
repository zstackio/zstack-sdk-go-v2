// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryConsoleProxyAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestQueryConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryConsoleProxyAgent result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.ManagementIp, r.State)
	}
	golog.Infof("======================================")
}

func TestPageConsoleProxyAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestPageConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("PageConsoleProxyAgent result: total=%d, returned=%d", total, len(result))
}

func TestGetConsoleProxyAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestGetConsoleProxyAgent Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ConsoleProxyAgent found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetConsoleProxyAgent(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("GetConsoleProxyAgent result: %s, IP: %s", result.UUID, result.ManagementIp)
}
