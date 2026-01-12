// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAgentVersion(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAgentVersion(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAgentVersion error: %v", err)
		return
	}
	golog.Infof("QueryAgentVersion result count: %d", len(result))
}
