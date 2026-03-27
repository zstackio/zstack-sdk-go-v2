// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTwoFactorAuthentication(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTwoFactorAuthentication(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryTwoFactorAuthentication error: %v", err)
		return
	}
	golog.Infof("QueryTwoFactorAuthentication result count: %d", len(result))
}

