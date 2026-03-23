// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryStackTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryStackTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryStackTemplate error: %v", err)
		return
	}
	golog.Infof("QueryStackTemplate result count: %d", len(result))
}

