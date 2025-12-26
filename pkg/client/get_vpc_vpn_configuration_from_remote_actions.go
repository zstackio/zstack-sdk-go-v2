// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcVpnConfigurationFromRemote gets VpcVpnConfigurationFromRemote by uuid
func (cli *ZSClient) GetVpcVpnConfigurationFromRemote(uuid string) (*view.GetVpcVpnConfigurationFromRemoteView, error) {
	var resp view.GetVpcVpnConfigurationFromRemoteView
	if err := cli.Get("v1/hybrid/vpn-conf/{uuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
