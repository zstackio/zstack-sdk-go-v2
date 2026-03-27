// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSSnmpPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSSnmpPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSSnmpPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSSnmpPlatform result count: %d", len(result))
}

