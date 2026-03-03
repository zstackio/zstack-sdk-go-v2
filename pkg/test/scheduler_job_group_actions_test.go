// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJobGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerJobGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJobGroup error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJobGroup result count: %d", len(result))
}

