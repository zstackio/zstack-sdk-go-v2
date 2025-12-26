// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedIpsec gets VpcAttachedIpsec by uuid
func (cli *ZSClient) GetVpcAttachedIpsec(uuid string) (*view.GetVpcAttachedIpsecView, error) {
	var resp view.GetVpcAttachedIpsecView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
