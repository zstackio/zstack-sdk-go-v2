// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceMdevDeviceSpecRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstanceMdevDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceMdevDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceMdevDeviceSpecRef result count: %d", len(result))
}

