// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryMonitorGroupEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupEventSubscription error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupEventSubscription result count: %d", len(result))
}

