// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateSriovPciDevices operates on GenerateSriovPciDevices
func (cli *ZSClient) GenerateSriovPciDevices(uuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
