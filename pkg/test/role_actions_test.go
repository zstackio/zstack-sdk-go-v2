// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryRole(&queryParam)
	if err != nil {
		t.Errorf("TestQueryRole error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryRole result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageRole(&queryParam)
	if err != nil {
		t.Errorf("TestPageRole error: %v", err)
		return
	}
	golog.Infof("PageRole result: total=%d, returned=%d", total, len(result))
}

func TestGetRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryRole(&queryParam)
	if err != nil {
		t.Errorf("TestGetRole Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Role found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetRole(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetRole error: %v", err)
		return
	}
	golog.Infof("GetRole result: %s, Name: %s", result.UUID, result.Name)
}
