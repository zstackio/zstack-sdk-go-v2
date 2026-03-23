// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNasFileSystem(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryNasFileSystem(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNasFileSystem error: %v", err)
		return
	}
	golog.Infof("QueryNasFileSystem result count: %d", len(result))
}

