// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmSchedHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmSchedHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedHistory error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedHistory result count: %d", len(result))
}

