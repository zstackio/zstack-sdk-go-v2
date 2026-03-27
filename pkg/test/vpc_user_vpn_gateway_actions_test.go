// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateVpcUserVpnGateway(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcUserVpnGatewayFromLocal(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateVpcUserVpnGateway Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcUserVpnGateway found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVpcUserVpnGatewayParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVpcUserVpnGatewayParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVpcUserVpnGateway(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVpcUserVpnGateway error: %v", err)
		return
	}
	golog.Infof("UpdateVpcUserVpnGateway result: %s", result.UUID)
}
