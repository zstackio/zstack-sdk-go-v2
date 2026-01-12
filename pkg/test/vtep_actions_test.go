// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVtep(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVtep(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVtep error: %v", err)
		return
	}
	golog.Infof("QueryVtep result count: %d", len(result))
}
