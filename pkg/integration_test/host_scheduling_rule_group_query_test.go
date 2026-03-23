// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostSchedulingRuleGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryHostSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("QueryHostSchedulingRuleGroup result count: %d", len(result))
}

