// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryContainerManagementEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryContainerManagementEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQueryContainerManagementEndpoint error: %v", err)
		return
	}
	golog.Infof("QueryContainerManagementEndpoint result count: %d", len(result))
}

