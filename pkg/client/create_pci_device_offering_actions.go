// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePciDeviceOffering creates PciDeviceOffering
func (cli *ZSClient) CreatePciDeviceOffering(params param.CreatePciDeviceOfferingParam) (*view.CreatePciDeviceOfferingEventView, error) {
	resp := view.CreatePciDeviceOfferingEventView{}
	if err := cli.Post("v1/pci-device/pci-device-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
