// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Bonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Bonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Bonding error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Bonding result count: %d", len(result))
}

