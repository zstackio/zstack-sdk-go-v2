// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterResourcePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterResourcePool(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterResourcePool error: %v", err)
		return
	}
	golog.Infof("QueryVCenterResourcePool result count: %d", len(result))
}

