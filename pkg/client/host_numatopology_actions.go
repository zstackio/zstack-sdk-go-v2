// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostNUMATopology 获取HostNUMATopology详情
func (cli *ZSClient) GetHostNUMATopology(uuid string) (*view.GetHostNUMATopologyEventView, error) {
	var resp view.GetHostNUMATopologyEventView
	if err := cli.Get("v1/hosts/{uuid}/numa", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

