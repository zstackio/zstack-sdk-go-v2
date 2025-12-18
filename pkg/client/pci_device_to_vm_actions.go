// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPciDeviceToVm 操作PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(params param.AttachPciDeviceToVmParam) (*view.AttachPciDeviceToVmEventView, error) {
	resp := view.AttachPciDeviceToVmEventView{}
	if err := cli.Post("v1/pci-device/pci-devices/{pciDeviceUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

