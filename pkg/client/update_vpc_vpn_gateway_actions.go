// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVpcVpnGateway updates VpcVpnGateway
func (cli *ZSClient) UpdateVpcVpnGateway(uuid string, params param.UpdateVpcVpnGatewayParam) (*view.UpdateVpcVpnGatewayEventView, error) {
	resp := view.UpdateVpcVpnGatewayEventView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
