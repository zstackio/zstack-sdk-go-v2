// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephOsdGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephOsdGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephOsdGroup error: %v", err)
		return
	}
	golog.Infof("QueryCephOsdGroup result count: %d", len(result))
}
