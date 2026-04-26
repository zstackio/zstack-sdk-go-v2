// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryResourceConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryResourceConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryResourceConfig error: %v", err)
		return
	}
	golog.Infof("QueryResourceConfig result count: %d", len(result))
}

