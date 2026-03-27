// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Instance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Instance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Instance error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Instance result count: %d", len(result))
}

