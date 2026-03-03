// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFlowMeter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryFlowMeter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFlowMeter error: %v", err)
		return
	}
	golog.Infof("QueryFlowMeter result count: %d", len(result))
}

