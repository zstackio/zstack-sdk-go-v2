// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryIAM2VirtualIDGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageIAM2VirtualIDGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("PageIAM2VirtualIDGroup result: total=%d, returned=%d", total, len(result))
}

func TestGetIAM2VirtualIDGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetIAM2VirtualIDGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("GetIAM2VirtualIDGroup result: %s, Name: %s", result.UUID, result.Name)
}
