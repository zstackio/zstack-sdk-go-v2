// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkInterfaceLldp(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkInterfaceLldp error: %v", err)
		return
	}
	golog.Infof("QueryHostNetworkInterfaceLldp result count: %d", len(result))
}

func TestGetHostNetworkInterfaceLldp(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostNetworkInterfaceLldp Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostNetworkInterfaceLldp found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetHostNetworkInterfaceLldp(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostNetworkInterfaceLldp error: %v", err)
		return
	}
	golog.Infof("GetHostNetworkInterfaceLldp result: %s", result.UUID)
}
