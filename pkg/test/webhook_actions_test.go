// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryWebhook(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestQueryWebhook error: %v", err)
		return
	}
	golog.Infof("QueryWebhook result count: %d", len(result))
}

