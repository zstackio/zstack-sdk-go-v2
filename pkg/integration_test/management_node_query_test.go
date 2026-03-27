// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryManagementNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryManagementNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryManagementNode error: %v", err)
		return
	}
	golog.Infof("QueryManagementNode result count: %d", len(result))
}

