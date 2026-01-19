// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicy error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryPolicy result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPagePolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PagePolicy(&queryParam)
	if err != nil {
		t.Errorf("TestPagePolicy error: %v", err)
		return
	}
	golog.Infof("PagePolicy result: total=%d, returned=%d", total, len(result))
}

func TestGetPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestGetPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Policy found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetPolicy(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPolicy error: %v", err)
		return
	}
	golog.Infof("GetPolicy result: %s, Name: %s", result.UUID, result.Name)
}
