// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVpcVpnConnectionRemote creates VpcVpnConnectionRemote
func (cli *ZSClient) CreateVpcVpnConnectionRemote(params param.CreateVpcVpnConnectionRemoteParam) (*view.CreateVpcVpnConnectionRemoteEventView, error) {
	resp := view.CreateVpcVpnConnectionRemoteEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
