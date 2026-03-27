// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStoragePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCephPrimaryStoragePool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStoragePool error: %v", err)
		return
	}
	golog.Infof("QueryCephPrimaryStoragePool result count: %d", len(result))
}

