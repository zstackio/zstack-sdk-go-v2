// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmCdRom(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmCdRom(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmCdRom error: %v", err)
		return
	}
	golog.Infof("QueryVmCdRom result count: %d", len(result))
}

