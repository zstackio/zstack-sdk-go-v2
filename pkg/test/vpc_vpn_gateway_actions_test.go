// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateVpcVpnGateway(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcVpnGatewayFromLocal(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVpcVpnGateway Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcVpnGateway found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVpcVpnGatewayParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVpcVpnGatewayParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVpcVpnGateway(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVpcVpnGateway error: %v", err)
		return
	}
	golog.Infof("UpdateVpcVpnGateway result: %s", result.UUID)
}
