// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAttachableVpcL3Network gets AttachableVpcL3Network by uuid
func (cli *ZSClient) GetAttachableVpcL3Network(uuid string) (*view.GetAttachableVpcL3NetworkView, error) {
	var resp view.GetAttachableVpcL3NetworkView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attachable-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
