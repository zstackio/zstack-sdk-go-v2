// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCdpPolicy error: %v", err)
		return
	}
	golog.Infof("QueryCdpPolicy result count: %d", len(result))
}

