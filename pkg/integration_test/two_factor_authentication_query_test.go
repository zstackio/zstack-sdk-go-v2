// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTwoFactorAuthentication(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryTwoFactorAuthentication(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTwoFactorAuthentication error: %v", err)
		return
	}
	golog.Infof("QueryTwoFactorAuthentication result count: %d", len(result))
}

