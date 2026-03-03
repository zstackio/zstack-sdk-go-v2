// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryConsoleProxyAgent(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryConsoleProxyAgent(&queryParam)
	if err != nil {
		t.Errorf("TestQueryConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("QueryConsoleProxyAgent result count: %d", len(result))
}

