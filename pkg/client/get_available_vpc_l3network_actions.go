// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAvailableVpcL3Network gets AvailableVpcL3Network by uuid
func (cli *ZSClient) GetAvailableVpcL3Network(uuid string) (*view.GetAvailableVpcL3NetworkView, error) {
	var resp view.GetAvailableVpcL3NetworkView
	if err := cli.Get("v1/vpc/virtual-routers/available-vpc-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
