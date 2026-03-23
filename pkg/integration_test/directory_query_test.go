// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDirectory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryDirectory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDirectory error: %v", err)
		return
	}
	golog.Infof("QueryDirectory result count: %d", len(result))
}

