// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJobHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySchedulerJobHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJobHistory error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJobHistory result count: %d", len(result))
}

