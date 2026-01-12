// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFcHbaDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFcHbaDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFcHbaDevice error: %v", err)
		return
	}
	golog.Infof("QueryFcHbaDevice result count: %d", len(result))
}
