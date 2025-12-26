// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(uuid string, params param.UpdatePciDeviceParam) (*view.UpdatePciDeviceEventView, error) {
	resp := view.UpdatePciDeviceEventView{}
	if err := cli.Put("v1/pci-device/pci-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
