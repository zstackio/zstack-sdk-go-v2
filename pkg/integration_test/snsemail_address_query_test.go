// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailAddress(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSEmailAddress(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailAddress error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailAddress result count: %d", len(result))
}

