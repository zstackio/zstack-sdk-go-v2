// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachPciDeviceFromVm operates on PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(params param.DetachPciDeviceFromVmParam) (*view.DetachPciDeviceFromVmEventView, error) {
	resp := view.DetachPciDeviceFromVmEventView{}
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
