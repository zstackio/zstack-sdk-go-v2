// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccount(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccount error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAccount result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageAccount(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAccount(&queryParam)
	if err != nil {
		t.Errorf("TestPageAccount error: %v", err)
		return
	}
	golog.Infof("PageAccount result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Type)
	}
}

func TestGetAccount(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccount Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Account found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAccount(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccount error: %v", err)
		return
	}
	golog.Infof("GetAccount result: %s, Name: %s", result.UUID, result.Name)
}
