// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImagePackage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImagePackage(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryImagePackage error: %v", err)
		return
	}
	golog.Infof("QueryImagePackage result count: %d", len(result))
}

