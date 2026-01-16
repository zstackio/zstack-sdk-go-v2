// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessKey error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAccessKey result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.AccessKeyID, r.State)
	}
	golog.Infof("======================================")
}

func TestPageAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestPageAccessKey error: %v", err)
		return
	}
	golog.Infof("PageAccessKey result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.AccessKeyID, r.State)
	}
}

func TestGetAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAccessKey(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccessKey error: %v", err)
		return
	}
	golog.Infof("GetAccessKey result: %s, AccessKeyId: %s", result.UUID, result.AccessKeyID)
}
