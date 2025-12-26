// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVpcUserVpnGateway updates VpcUserVpnGateway
func (cli *ZSClient) UpdateVpcUserVpnGateway(uuid string, params param.UpdateVpcUserVpnGatewayParam) (*view.UpdateVpcUserVpnGatewayEventView, error) {
	resp := view.UpdateVpcUserVpnGatewayEventView{}
	if err := cli.Put("v1/hybrid/user-vpn/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
