// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryBackupStorage result count: %d", len(result))
}

