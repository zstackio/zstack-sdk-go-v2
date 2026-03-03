// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryExponBlockVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryExponBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryExponBlockVolume error: %v", err)
		return
	}
	golog.Infof("QueryExponBlockVolume result count: %d", len(result))
}

