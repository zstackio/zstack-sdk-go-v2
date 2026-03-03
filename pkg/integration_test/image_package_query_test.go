// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImagePackage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryImagePackage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImagePackage error: %v", err)
		return
	}
	golog.Infof("QueryImagePackage result count: %d", len(result))
}

