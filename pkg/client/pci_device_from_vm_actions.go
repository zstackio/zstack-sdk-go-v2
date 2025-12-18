// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPciDeviceFromVm 操作PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(params param.DetachPciDeviceFromVmParam) (*view.DetachPciDeviceFromVmEventView, error) {
	resp := view.DetachPciDeviceFromVmEventView{}
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

