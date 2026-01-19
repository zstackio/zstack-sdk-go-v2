// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedResource(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedResource error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySharedResource result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.OwnerAccountUuid, r.ResourceType)
	}
	golog.Infof("======================================")
}

func TestPageSharedResource(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestPageSharedResource error: %v", err)
		return
	}
	golog.Infof("PageSharedResource result: total=%d, returned=%d", total, len(result))
}
