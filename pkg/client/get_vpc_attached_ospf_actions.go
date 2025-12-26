// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedOspf gets VpcAttachedOspf by uuid
func (cli *ZSClient) GetVpcAttachedOspf(uuid string) (*view.GetVpcAttachedOspfView, error) {
	var resp view.GetVpcAttachedOspfView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-ospf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
