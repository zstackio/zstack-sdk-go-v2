// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeLun error: %v", err)
		return
	}
	golog.Infof("QueryNvmeLun result count: %d", len(result))
}
