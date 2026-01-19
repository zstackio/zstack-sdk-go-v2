// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCdpPolicy error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryCdpPolicy result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestPageCdpPolicy error: %v", err)
		return
	}
	golog.Infof("PageCdpPolicy result: total=%d, returned=%d", total, len(result))
}

func TestGetCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestGetCdpPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpPolicy found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetCdpPolicy(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCdpPolicy error: %v", err)
		return
	}
	golog.Infof("GetCdpPolicy result: %s, Name: %s", result.UUID, result.Name)
}
