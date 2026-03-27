// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroup error: %v", err)
		return
	}
	golog.Infof("QuerySecurityGroup result count: %d", len(result))
}

