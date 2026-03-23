// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryExternalBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryExternalBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryExternalBackup error: %v", err)
		return
	}
	golog.Infof("QueryExternalBackup result count: %d", len(result))
}

