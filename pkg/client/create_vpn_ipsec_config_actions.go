// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpnIpsecConfig creates VpnIpsecConfig
func (cli *ZSClient) CreateVpnIpsecConfig(params param.CreateVpnIpsecConfigParam) (*view.CreateVpnIpsecConfigEventView, error) {
	resp := view.CreateVpnIpsecConfigEventView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
