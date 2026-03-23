// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlarm error: %v", err)
		return
	}
	golog.Infof("QueryAlarm result count: %d", len(result))
}

