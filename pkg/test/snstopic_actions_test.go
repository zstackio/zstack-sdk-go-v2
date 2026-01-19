// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSTopic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTopic error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySNSTopic result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageSNSTopic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestPageSNSTopic error: %v", err)
		return
	}
	golog.Infof("PageSNSTopic result: total=%d, returned=%d", total, len(result))
}

func TestGetSNSTopic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSTopic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTopic found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSNSTopic(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSTopic error: %v", err)
		return
	}
	golog.Infof("GetSNSTopic result: %s, Name: %s", result.UUID, result.Name)
}
