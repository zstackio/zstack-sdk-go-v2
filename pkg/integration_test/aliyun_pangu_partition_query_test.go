// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAliyunPanguPartition(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAliyunPanguPartition(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAliyunPanguPartition error: %v", err)
		return
	}
	golog.Infof("QueryAliyunPanguPartition result count: %d", len(result))
}

