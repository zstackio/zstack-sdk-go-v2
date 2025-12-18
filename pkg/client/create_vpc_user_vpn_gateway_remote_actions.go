// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcUserVpnGatewayRemote creates VpcUserVpnGatewayRemote
func (cli *ZSClient) CreateVpcUserVpnGatewayRemote(params param.CreateVpcUserVpnGatewayRemoteParam) (*view.CreateVpcUserVpnGatewayRemoteEventView, error) {
	resp := view.CreateVpcUserVpnGatewayRemoteEventView{}
	if err := cli.Post("v1/hybrid/user-vpn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
