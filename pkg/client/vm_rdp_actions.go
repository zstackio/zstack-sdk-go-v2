// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmRDP 获取VmRDP详情
func (cli *ZSClient) GetVmRDP(uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.Get("v1/vm-instances/{uuid}/rdp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

