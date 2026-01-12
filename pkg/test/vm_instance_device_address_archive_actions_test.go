// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceDeviceAddressArchive(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstanceDeviceAddressArchive(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceDeviceAddressArchive error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceDeviceAddressArchive result count: %d", len(result))
}
func TestGetVmInstanceDeviceAddressArchive(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmInstanceDeviceAddressArchive(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstanceDeviceAddressArchive Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstanceDeviceAddressArchive found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmInstanceDeviceAddressArchive(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmInstanceDeviceAddressArchive error: %v", err)
		return
	}
	golog.Infof("GetVmInstanceDeviceAddressArchive result: %s", result.UUID)
}
