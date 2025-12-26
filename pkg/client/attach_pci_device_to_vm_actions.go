// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(params param.AttachPciDeviceToVmParam) (*view.AttachPciDeviceToVmEventView, error) {
	resp := view.AttachPciDeviceToVmEventView{}
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
