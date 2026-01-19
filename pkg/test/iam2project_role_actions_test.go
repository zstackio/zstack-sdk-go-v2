// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryIAM2ProjectRole(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectRole error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryIAM2ProjectRole result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageIAM2ProjectRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageIAM2ProjectRole(&queryParam)
	if err != nil {
		t.Errorf("TestPageIAM2ProjectRole error: %v", err)
		return
	}
	golog.Infof("PageIAM2ProjectRole result: total=%d, returned=%d", total, len(result))
}

func TestGetIAM2ProjectRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryIAM2ProjectRole(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectRole Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectRole found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetIAM2ProjectRole(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectRole error: %v", err)
		return
	}
	golog.Infof("GetIAM2ProjectRole result: %s, Name: %s", result.UUID, result.Name)
}
