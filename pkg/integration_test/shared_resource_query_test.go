// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedResource(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedResource error: %v", err)
		return
	}
	golog.Infof("QuerySharedResource result count: %d", len(result))
}

