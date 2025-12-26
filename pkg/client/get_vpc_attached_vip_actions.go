// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(uuid string) (*view.GetVpcAttachedVipView, error) {
	var resp view.GetVpcAttachedVipView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-vip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
