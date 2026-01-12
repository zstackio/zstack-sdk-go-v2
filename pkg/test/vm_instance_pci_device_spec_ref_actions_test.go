// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstancePciDeviceSpecRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstancePciDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstancePciDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstancePciDeviceSpecRef result count: %d", len(result))
}
func TestGetVmInstancePciDeviceSpecRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmInstancePciDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstancePciDeviceSpecRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstancePciDeviceSpecRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmInstancePciDeviceSpecRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmInstancePciDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("GetVmInstancePciDeviceSpecRef result: %s", result.UUID)
}
