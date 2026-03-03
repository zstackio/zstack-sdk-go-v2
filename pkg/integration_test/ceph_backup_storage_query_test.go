// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephBackupStorage result count: %d", len(result))
}

