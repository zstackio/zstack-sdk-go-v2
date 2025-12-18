// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmvNUMATopology 获取VmvNUMATopology详情
func (cli *ZSClient) GetVmvNUMATopology(uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.Get("v1/vm-instances/{uuid}/vnuma-topology", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

