// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSApplicationPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSApplicationPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSApplicationPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSApplicationPlatform result count: %d", len(result))
}

