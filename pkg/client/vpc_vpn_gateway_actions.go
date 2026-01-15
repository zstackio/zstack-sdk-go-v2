// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcVpnGateway updates VpcVpnGateway
func (cli *ZSClient) UpdateVpcVpnGateway(uuid string, params param.UpdateVpcVpnGatewayParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.Put("v1/hybrid/vpc-vpn", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
