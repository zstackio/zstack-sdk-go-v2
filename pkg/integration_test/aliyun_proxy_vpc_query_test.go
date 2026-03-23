// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAliyunProxyVpc(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAliyunProxyVpc(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAliyunProxyVpc error: %v", err)
		return
	}
	golog.Infof("QueryAliyunProxyVpc result count: %d", len(result))
}

