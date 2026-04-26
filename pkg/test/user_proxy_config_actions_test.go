// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserProxyConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUserProxyConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserProxyConfig error: %v", err)
		return
	}
	golog.Infof("QueryUserProxyConfig result count: %d", len(result))
}

