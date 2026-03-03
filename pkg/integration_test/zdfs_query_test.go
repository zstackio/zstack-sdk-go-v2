// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZdfs(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryZdfs(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZdfs error: %v", err)
		return
	}
	golog.Infof("QueryZdfs result count: %d", len(result))
}

