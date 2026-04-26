// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelCenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelCenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelCenter error: %v", err)
		return
	}
	golog.Infof("QueryModelCenter result count: %d", len(result))
}

