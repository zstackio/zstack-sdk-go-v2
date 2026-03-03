// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBlockPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBlockPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBlockPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryBlockPrimaryStorage result count: %d", len(result))
}

