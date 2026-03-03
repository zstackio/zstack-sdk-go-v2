// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostPhysicalMemory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryHostPhysicalMemory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostPhysicalMemory error: %v", err)
		return
	}
	golog.Infof("QueryHostPhysicalMemory result count: %d", len(result))
}

