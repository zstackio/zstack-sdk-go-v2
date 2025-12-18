// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcAttachedNetflow 获取VpcAttachedNetflow详情
func (cli *ZSClient) GetVpcAttachedNetflow(uuid string) (*view.GetVpcAttachedNetflowView, error) {
	var resp view.GetVpcAttachedNetflowView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-netflow", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

