// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAttachableVpcL3Network 获取AttachableVpcL3Network详情
func (cli *ZSClient) GetAttachableVpcL3Network(uuid string) (*view.GetAttachableVpcL3NetworkView, error) {
	var resp view.GetAttachableVpcL3NetworkView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attachable-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

