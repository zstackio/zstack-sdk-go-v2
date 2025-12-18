// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVipUsedPorts gets VipUsedPorts by uuid
func (cli *ZSClient) GetVipUsedPorts(uuid string) (*view.GetVipUsedPortsView, error) {
	var resp view.GetVipUsedPortsView
	if err := cli.Get("v1/vips/{uuid}/usedports", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
