// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcAttachedOspf 获取VpcAttachedOspf详情
func (cli *ZSClient) GetVpcAttachedOspf(uuid string) (*view.GetVpcAttachedOspfView, error) {
	var resp view.GetVpcAttachedOspfView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-ospf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

