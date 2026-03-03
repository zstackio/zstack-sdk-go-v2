// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEip(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEip error: %v", err)
		return
	}
	golog.Infof("QueryEip result count: %d", len(result))
}

