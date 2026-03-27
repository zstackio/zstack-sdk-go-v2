// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcUserVpnGateway updates VpcUserVpnGateway
func (cli *ZSClient) UpdateVpcUserVpnGateway(ctx context.Context, uuid string, params param.UpdateVpcUserVpnGatewayParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/user-vpn", uuid, "", map[string]interface{}{
		"updateVpcUserVpnGateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
