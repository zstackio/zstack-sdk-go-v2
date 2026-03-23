// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAliyunProxyVSwitch(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAliyunProxyVSwitch(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAliyunProxyVSwitch error: %v", err)
		return
	}
	golog.Infof("QueryAliyunProxyVSwitch result count: %d", len(result))
}

