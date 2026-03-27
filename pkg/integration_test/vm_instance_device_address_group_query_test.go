// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceDeviceAddressGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstanceDeviceAddressGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceDeviceAddressGroup error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceDeviceAddressGroup result count: %d", len(result))
}

