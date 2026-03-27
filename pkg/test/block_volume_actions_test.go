// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBlockVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBlockVolume(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryBlockVolume error: %v", err)
		return
	}
	golog.Infof("QueryBlockVolume result count: %d", len(result))
}

