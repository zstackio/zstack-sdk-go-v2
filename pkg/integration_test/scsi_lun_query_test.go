// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryScsiLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryScsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryScsiLun error: %v", err)
		return
	}
	golog.Infof("QueryScsiLun result count: %d", len(result))
}

