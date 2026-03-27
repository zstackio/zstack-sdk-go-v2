// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshot(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("QueryVolumeSnapshot result count: %d", len(result))
}

