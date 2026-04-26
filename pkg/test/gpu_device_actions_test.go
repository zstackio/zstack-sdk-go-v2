// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGpuDevice error: %v", err)
		return
	}
	golog.Infof("QueryGpuDevice result count: %d", len(result))
}

