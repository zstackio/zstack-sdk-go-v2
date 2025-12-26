// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmInstanceFirstBootDevice gets VmInstanceFirstBootDevice by uuid
func (cli *ZSClient) GetVmInstanceFirstBootDevice(uuid string) (*view.GetVmInstanceFirstBootDeviceView, error) {
	var resp view.GetVmInstanceFirstBootDeviceView
	if err := cli.Get("v1/vm-instances/{uuid}/first-boot-device", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
