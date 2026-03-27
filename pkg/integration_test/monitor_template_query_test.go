// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("QueryMonitorTemplate result count: %d", len(result))
}

