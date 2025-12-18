// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcVpnConfigurationFromRemote 获取VpcVpnConfigurationFromRemote详情
func (cli *ZSClient) GetVpcVpnConfigurationFromRemote(uuid string) (*view.GetVpcVpnConfigurationFromRemoteView, error) {
	var resp view.GetVpcVpnConfigurationFromRemoteView
	if err := cli.Get("v1/hybrid/vpn-conf/{uuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

