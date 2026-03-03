// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAliyunEbsBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAliyunEbsBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAliyunEbsBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryAliyunEbsBackupStorage result count: %d", len(result))
}

