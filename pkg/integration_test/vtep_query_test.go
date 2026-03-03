// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVtep(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVtep(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVtep error: %v", err)
		return
	}
	golog.Infof("QueryVtep result count: %d", len(result))
}

