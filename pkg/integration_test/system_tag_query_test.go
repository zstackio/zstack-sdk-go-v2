// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySystemTag(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySystemTag(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySystemTag error: %v", err)
		return
	}
	golog.Infof("QuerySystemTag result count: %d", len(result))
}

