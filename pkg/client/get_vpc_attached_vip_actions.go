// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(uuid string) (*view.GetVpcAttachedVipView, error) {
	var resp view.GetVpcAttachedVipView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-vip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
