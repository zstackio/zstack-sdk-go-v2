// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmDeviceAddress 获取VmDeviceAddress详情
func (cli *ZSClient) GetVmDeviceAddress(uuid string) (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.Get("v1/vm-instances/devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

