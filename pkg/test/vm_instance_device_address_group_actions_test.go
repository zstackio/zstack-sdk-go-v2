// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceDeviceAddressGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstanceDeviceAddressGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceDeviceAddressGroup error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceDeviceAddressGroup result count: %d", len(result))
}
func TestGetVmInstanceDeviceAddressGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmInstanceDeviceAddressGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstanceDeviceAddressGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstanceDeviceAddressGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmInstanceDeviceAddressGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmInstanceDeviceAddressGroup error: %v", err)
		return
	}
	golog.Infof("GetVmInstanceDeviceAddressGroup result: %s", result.UUID)
}
