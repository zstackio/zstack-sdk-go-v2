// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVpcVpnGateway 更新VpcVpnGateway
func (cli *ZSClient) UpdateVpcVpnGateway(uuid string, params param.UpdateVpcVpnGatewayParam) (*view.UpdateVpcVpnGatewayEventView, error) {
	resp := view.UpdateVpcVpnGatewayEventView{}
	if err := cli.Put("v1/hybrid/vpc-vpn/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

