// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2Organization(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2Organization error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryIAM2Organization result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageIAM2Organization(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestPageIAM2Organization error: %v", err)
		return
	}
	golog.Infof("PageIAM2Organization result: total=%d, returned=%d", total, len(result))
}

func TestGetIAM2Organization(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2Organization Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2Organization found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetIAM2Organization(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2Organization error: %v", err)
		return
	}
	golog.Infof("GetIAM2Organization result: %s, Name: %s", result.UUID, result.Name)
}
