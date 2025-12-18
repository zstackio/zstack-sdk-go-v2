// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmMonitorNumber 获取VmMonitorNumber详情
func (cli *ZSClient) GetVmMonitorNumber(uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.Get("v1/vm-instances/{uuid}/monitorNumber", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

