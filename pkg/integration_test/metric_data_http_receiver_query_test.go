// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMetricDataHttpReceiver(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryMetricDataHttpReceiver(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMetricDataHttpReceiver error: %v", err)
		return
	}
	golog.Infof("QueryMetricDataHttpReceiver result count: %d", len(result))
}

