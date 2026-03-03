// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBuildAppExportHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBuildAppExportHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBuildAppExportHistory error: %v", err)
		return
	}
	golog.Infof("QueryBuildAppExportHistory result count: %d", len(result))
}

