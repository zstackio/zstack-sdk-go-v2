// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmDeviceAddress gets VmDeviceAddress by uuid
func (cli *ZSClient) GetVmDeviceAddress(uuid string) (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.Get("v1/vm-instances/devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
