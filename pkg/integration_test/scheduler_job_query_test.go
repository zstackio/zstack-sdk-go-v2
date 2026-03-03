// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySchedulerJob(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJob error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJob result count: %d", len(result))
}

