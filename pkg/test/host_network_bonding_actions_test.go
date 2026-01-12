// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("QueryHostNetworkBonding result count: %d", len(result))
}
