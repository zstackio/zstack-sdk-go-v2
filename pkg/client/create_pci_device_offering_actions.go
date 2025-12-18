// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePciDeviceOffering creates PciDeviceOffering
func (cli *ZSClient) CreatePciDeviceOffering(params param.CreatePciDeviceOfferingParam) (*view.CreatePciDeviceOfferingEventView, error) {
	resp := view.CreatePciDeviceOfferingEventView{}
	if err := cli.Post("v1/pci-device/pci-device-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
