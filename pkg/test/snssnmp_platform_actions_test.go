// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSSnmpPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSSnmpPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSSnmpPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSSnmpPlatform result count: %d", len(result))
}

