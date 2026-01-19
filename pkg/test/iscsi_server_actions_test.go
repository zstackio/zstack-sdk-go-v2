// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIscsiServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIscsiServer error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryIscsiServer result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageIscsiServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestPageIscsiServer error: %v", err)
		return
	}
	golog.Infof("PageIscsiServer result: total=%d, returned=%d", total, len(result))
}

func TestGetIscsiServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestGetIscsiServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IscsiServer found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetIscsiServer(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIscsiServer error: %v", err)
		return
	}
	golog.Infof("GetIscsiServer result: %s, Name: %s", result.UUID, result.Name)
}
