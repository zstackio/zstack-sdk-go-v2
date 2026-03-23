// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetworkPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("QueryL2VxlanNetworkPool result count: %d", len(result))
}

