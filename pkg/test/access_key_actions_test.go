// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessKey error: %v", err)
		return
	}
	golog.Infof("QueryAccessKey result count: %d", len(result))
}

