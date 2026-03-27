// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2Network(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2Network(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2Network error: %v", err)
		return
	}
	golog.Infof("QueryL2Network result count: %d", len(result))
}

