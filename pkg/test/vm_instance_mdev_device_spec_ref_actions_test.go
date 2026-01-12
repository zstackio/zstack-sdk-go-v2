// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceMdevDeviceSpecRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstanceMdevDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceMdevDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceMdevDeviceSpecRef result count: %d", len(result))
}
func TestGetVmInstanceMdevDeviceSpecRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmInstanceMdevDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmInstanceMdevDeviceSpecRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmInstanceMdevDeviceSpecRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmInstanceMdevDeviceSpecRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmInstanceMdevDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("GetVmInstanceMdevDeviceSpecRef result: %s", result.UUID)
}
