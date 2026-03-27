// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroup(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroup error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroup result count: %d", len(result))
}

