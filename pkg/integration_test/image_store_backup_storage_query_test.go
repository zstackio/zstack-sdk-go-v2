// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageStoreBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryImageStoreBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageStoreBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryImageStoreBackupStorage result count: %d", len(result))
}

