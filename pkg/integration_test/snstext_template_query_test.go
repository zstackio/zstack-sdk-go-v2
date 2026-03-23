// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSTextTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSTextTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTextTemplate error: %v", err)
		return
	}
	golog.Infof("QuerySNSTextTemplate result count: %d", len(result))
}

