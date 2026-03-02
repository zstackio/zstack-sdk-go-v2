// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIscsiServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIscsiServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIscsiServer error: %v", err)
		return
	}
	golog.Infof("QueryIscsiServer result count: %d", len(result))
}

