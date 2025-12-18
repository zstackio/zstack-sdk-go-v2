// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAvailableVpcL3Network 获取AvailableVpcL3Network详情
func (cli *ZSClient) GetAvailableVpcL3Network(uuid string) (*view.GetAvailableVpcL3NetworkView, error) {
	var resp view.GetAvailableVpcL3NetworkView
	if err := cli.Get("v1/vpc/virtual-routers/available-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

