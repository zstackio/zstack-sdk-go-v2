// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(params param.AttachPciDeviceToVmParam) (*view.AttachPciDeviceToVmEventView, error) {
	resp := view.AttachPciDeviceToVmEventView{}
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
