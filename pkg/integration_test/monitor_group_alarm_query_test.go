// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryMonitorGroupAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupAlarm error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupAlarm result count: %d", len(result))
}

