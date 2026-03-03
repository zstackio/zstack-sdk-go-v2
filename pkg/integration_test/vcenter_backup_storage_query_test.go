// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVCenterBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryVCenterBackupStorage result count: %d", len(result))
}

