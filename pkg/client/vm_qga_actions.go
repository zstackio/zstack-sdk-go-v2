// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmQga 获取VmQga详情
func (cli *ZSClient) GetVmQga(uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.Get("v1/vm-instances/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

