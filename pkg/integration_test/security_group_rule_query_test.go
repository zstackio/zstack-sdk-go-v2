// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroupRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("QuerySecurityGroupRule result count: %d", len(result))
}

