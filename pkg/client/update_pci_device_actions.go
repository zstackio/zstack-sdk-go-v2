// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(uuid string, params param.UpdatePciDeviceParam) (*view.UpdatePciDeviceEventView, error) {
	resp := view.UpdatePciDeviceEventView{}
	if err := cli.Put("v1/pci-device/pci-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
