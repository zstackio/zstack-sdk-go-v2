// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmSchedulingRuleGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedulingRuleGroup result count: %d", len(result))
}

