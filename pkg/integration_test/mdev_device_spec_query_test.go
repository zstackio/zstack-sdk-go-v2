// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMdevDeviceSpec(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryMdevDeviceSpec(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMdevDeviceSpec error: %v", err)
		return
	}
	golog.Infof("QueryMdevDeviceSpec result count: %d", len(result))
}

