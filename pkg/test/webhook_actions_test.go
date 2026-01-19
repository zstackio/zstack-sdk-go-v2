// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryWebhook(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestQueryWebhook error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryWebhook result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageWebhook(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestPageWebhook error: %v", err)
		return
	}
	golog.Infof("PageWebhook result: total=%d, returned=%d", total, len(result))
}

func TestGetWebhook(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestGetWebhook Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Webhook found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetWebhook(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetWebhook error: %v", err)
		return
	}
	golog.Infof("GetWebhook result: %s, Name: %s", result.UUID, result.Name)
}
