// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVpnIpsecConfig creates VpnIpsecConfig
func (cli *ZSClient) CreateVpnIpsecConfig(params param.CreateVpnIpsecConfigParam) (*view.CreateVpnIpsecConfigEventView, error) {
	resp := view.CreateVpnIpsecConfigEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
