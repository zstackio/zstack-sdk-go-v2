// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerTrigger(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerTrigger result count: %d", len(result))
}

