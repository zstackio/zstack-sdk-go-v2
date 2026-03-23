// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryVmInstance result count: %d", len(result))
}

