// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJobHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerJobHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJobHistory error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJobHistory result count: %d", len(result))
}
