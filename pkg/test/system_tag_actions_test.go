// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySystemTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySystemTag error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySystemTag result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Tag, r.ResourceType)
	}
	golog.Infof("======================================")
}

func TestPageSystemTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestPageSystemTag error: %v", err)
		return
	}
	golog.Infof("PageSystemTag result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Tag, r.ResourceType)
	}
}

func TestGetSystemTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestGetSystemTag Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SystemTag found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSystemTag(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSystemTag error: %v", err)
		return
	}
	golog.Infof("GetSystemTag result: %s, Tag: %s", result.UUID, result.Tag)
}
