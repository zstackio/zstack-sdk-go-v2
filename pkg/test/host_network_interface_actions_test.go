// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkInterface(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkInterface error: %v", err)
		return
	}
	golog.Infof("QueryHostNetworkInterface result count: %d", len(result))
}

func TestUpdateHostNetworkInterface(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateHostNetworkInterface Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostNetworkInterface found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateHostNetworkInterfaceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateHostNetworkInterfaceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateHostNetworkInterface(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateHostNetworkInterface error: %v", err)
		return
	}
	golog.Infof("UpdateHostNetworkInterface result: %s", result.UUID)
}

func TestLocateHostNetworkInterface(t *testing.T) {
	// LocateHostNetworkInterface operation
	t.Skip("TestLocateHostNetworkInterface requires manual implementation")

}
