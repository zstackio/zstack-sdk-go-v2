// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSDingTalkEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySNSDingTalkEndpoint result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageSNSDingTalkEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestPageSNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("PageSNSDingTalkEndpoint result: total=%d, returned=%d", total, len(result))
}

func TestGetSNSDingTalkEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSDingTalkEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSDingTalkEndpoint found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSNSDingTalkEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("GetSNSDingTalkEndpoint result: %s, Name: %s", result.UUID, result.Name)
}
