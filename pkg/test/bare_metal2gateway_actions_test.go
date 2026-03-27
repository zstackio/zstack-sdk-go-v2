// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Gateway(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Gateway(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Gateway error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Gateway result count: %d", len(result))
}

