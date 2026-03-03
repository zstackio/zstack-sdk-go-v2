// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryNvmeServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeServer error: %v", err)
		return
	}
	golog.Infof("QueryNvmeServer result count: %d", len(result))
}

