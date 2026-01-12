// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccountResourceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccountResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccountResourceRef error: %v", err)
		return
	}
	golog.Infof("QueryAccountResourceRef result count: %d", len(result))
}
