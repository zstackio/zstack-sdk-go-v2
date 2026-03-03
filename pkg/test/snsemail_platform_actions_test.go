// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEmailPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailPlatform result count: %d", len(result))
}

