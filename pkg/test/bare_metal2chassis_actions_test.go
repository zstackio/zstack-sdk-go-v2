// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Chassis(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Chassis(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Chassis error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Chassis result count: %d", len(result))
}

