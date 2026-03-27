// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVCenterPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryVCenterPrimaryStorage result count: %d", len(result))
}

