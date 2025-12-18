// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcVpnConnectionRemote 创建VpcVpnConnectionRemote
func (cli *ZSClient) CreateVpcVpnConnectionRemote(params param.CreateVpcVpnConnectionRemoteParam) (*view.CreateVpcVpnConnectionRemoteEventView, error) {
	resp := view.CreateVpcVpnConnectionRemoteEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

