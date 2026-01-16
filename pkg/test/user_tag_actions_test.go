// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryUserTag(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserTag error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryUserTag result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Tag, r.ResourceType)
	}
	golog.Infof("======================================")
}

func TestPageUserTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageUserTag(&queryParam)
	if err != nil {
		t.Errorf("TestPageUserTag error: %v", err)
		return
	}
	golog.Infof("PageUserTag result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Tag, r.ResourceType)
	}
}

func TestGetUserTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryUserTag(&queryParam)
	if err != nil {
		t.Errorf("TestGetUserTag Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserTag found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetUserTag(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUserTag error: %v", err)
		return
	}
	golog.Infof("GetUserTag result: %s, Tag: %s", result.UUID, result.Tag)
}
