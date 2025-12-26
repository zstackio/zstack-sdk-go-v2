// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UngenerateSriovPciDevices operates on UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(uuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
