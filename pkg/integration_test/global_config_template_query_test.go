// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGlobalConfigTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryGlobalConfigTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfigTemplate error: %v", err)
		return
	}
	golog.Infof("QueryGlobalConfigTemplate result count: %d", len(result))
}

