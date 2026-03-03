// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCbtTask(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCbtTask(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCbtTask error: %v", err)
		return
	}
	golog.Infof("QueryCbtTask result count: %d", len(result))
}

