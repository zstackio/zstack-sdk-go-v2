// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUser(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUser error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryUser result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageUser(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageUser(&queryParam)
	if err != nil {
		t.Errorf("TestPageUser error: %v", err)
		return
	}
	golog.Infof("PageUser result: total=%d, returned=%d", total, len(result))
}

func TestGetUser(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestGetUser Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No User found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetUser(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUser error: %v", err)
		return
	}
	golog.Infof("GetUser result: %s, Name: %s", result.UUID, result.Name)
}
