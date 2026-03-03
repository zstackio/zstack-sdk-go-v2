// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorTrigger error: %v", err)
		return
	}
	golog.Infof("QueryMonitorTrigger result count: %d", len(result))
}

