// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEventSubscription error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryEventSubscription result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestPageEventSubscription error: %v", err)
		return
	}
	golog.Infof("PageEventSubscription result: total=%d, returned=%d", total, len(result))
}

func TestGetEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestGetEventSubscription Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventSubscription found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetEventSubscription(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetEventSubscription error: %v", err)
		return
	}
	golog.Infof("GetEventSubscription result: %s, Name: %s", result.UUID, result.Name)
}
