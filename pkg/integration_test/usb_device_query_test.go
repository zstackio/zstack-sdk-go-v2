// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUsbDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryUsbDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUsbDevice error: %v", err)
		return
	}
	golog.Infof("QueryUsbDevice result count: %d", len(result))
}

