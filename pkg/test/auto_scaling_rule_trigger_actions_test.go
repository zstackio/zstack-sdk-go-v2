// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingRuleTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingRuleTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingRuleTrigger error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingRuleTrigger result count: %d", len(result))
}

