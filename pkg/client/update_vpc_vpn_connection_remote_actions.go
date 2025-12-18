// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVpcVpnConnectionRemote updates VpcVpnConnectionRemote
func (cli *ZSClient) UpdateVpcVpnConnectionRemote(uuid string, params param.UpdateVpcVpnConnectionRemoteParam) (*view.UpdateVpcVpnConnectionRemoteEventView, error) {
	resp := view.UpdateVpcVpnConnectionRemoteEventView{}
	if err := cli.Put("v1/hybrid/vpn-connection/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
