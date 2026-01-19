// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfig error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryGlobalConfig result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.Category, r.Name, r.Value)
	}
	golog.Infof("======================================")
}

func TestPageGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestPageGlobalConfig error: %v", err)
		return
	}
	golog.Infof("PageGlobalConfig result: total=%d, returned=%d", total, len(result))
}
