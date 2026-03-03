// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlock(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySharedBlock(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlock error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlock result count: %d", len(result))
}

