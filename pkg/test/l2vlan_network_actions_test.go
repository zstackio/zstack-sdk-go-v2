// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VlanNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VlanNetwork result count: %d", len(result))
}

