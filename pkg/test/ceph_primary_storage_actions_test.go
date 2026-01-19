// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryCephPrimaryStorage result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestPageCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("PageCephPrimaryStorage result: total=%d, returned=%d", total, len(result))
}

func TestGetCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStorage found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetCephPrimaryStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("GetCephPrimaryStorage result: %s, Name: %s", result.UUID, result.Name)
}
