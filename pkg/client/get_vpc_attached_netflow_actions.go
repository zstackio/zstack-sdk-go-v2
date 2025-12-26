// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedNetflow gets VpcAttachedNetflow by uuid
func (cli *ZSClient) GetVpcAttachedNetflow(uuid string) (*view.GetVpcAttachedNetflowView, error) {
	var resp view.GetVpcAttachedNetflowView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-netflow", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
