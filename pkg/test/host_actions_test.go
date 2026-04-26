// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHost error: %v", err)
		return
	}
	golog.Infof("QueryHost result count: %d", len(result))
}

