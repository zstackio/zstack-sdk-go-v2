// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSDingTalkEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSDingTalkEndpoint result count: %d", len(result))
}

