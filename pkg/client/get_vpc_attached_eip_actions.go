// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcAttachedEip gets VpcAttachedEip by uuid
func (cli *ZSClient) GetVpcAttachedEip(uuid string) (*view.GetVpcAttachedEipView, error) {
	var resp view.GetVpcAttachedEipView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-eip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
