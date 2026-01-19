// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryQuota(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryQuota(&queryParam)
	if err != nil {
		t.Errorf("TestQueryQuota error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryQuota result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d", r.Name, r.IdentityType, r.Value)
	}
	golog.Infof("======================================")
}

func TestPageQuota(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageQuota(&queryParam)
	if err != nil {
		t.Errorf("TestPageQuota error: %v", err)
		return
	}
	golog.Infof("PageQuota result: total=%d, returned=%d", total, len(result))
}
