// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNativeHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNativeHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNativeHost error: %v", err)
		return
	}
	golog.Infof("QueryNativeHost result count: %d", len(result))
}
