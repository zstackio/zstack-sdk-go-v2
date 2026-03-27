// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryExternalServiceConfiguration(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryExternalServiceConfiguration(&queryParam)
	if err != nil {
		t.Errorf("TestQueryExternalServiceConfiguration error: %v", err)
		return
	}
	golog.Infof("QueryExternalServiceConfiguration result count: %d", len(result))
}

