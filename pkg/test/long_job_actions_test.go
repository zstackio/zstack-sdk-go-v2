// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLongJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLongJob error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryLongJob result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageLongJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestPageLongJob error: %v", err)
		return
	}
	golog.Infof("PageLongJob result: total=%d, returned=%d", total, len(result))
}

func TestGetLongJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestGetLongJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LongJob found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetLongJob(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLongJob error: %v", err)
		return
	}
	golog.Infof("GetLongJob result: %s, Name: %s", result.UUID, result.Name)
}
