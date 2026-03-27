// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcSnatState(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcSnatState(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcSnatState error: %v", err)
		return
	}
	golog.Infof("QueryVpcSnatState result count: %d", len(result))
}

