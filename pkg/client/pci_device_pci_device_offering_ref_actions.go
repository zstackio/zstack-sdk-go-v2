// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryPciDevicePciDeviceOffering queries PciDevicePciDeviceOffering list
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp []view.PciDevicePciDeviceOfferingRefInventoryView
	return resp, cli.List("v1/pci-devices/pci-devices/pci-device-offerings", params, &resp)
}

func (cli *ZSClient) GetPciDevicePciDeviceOffering(uuid string) (*view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp view.PciDevicePciDeviceOfferingRefInventoryView
	if err := cli.Get("v1/pci-devices/pci-devices/pci-device-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDevicePciDeviceOffering Pagination
func (cli *ZSClient) PagePciDevicePciDeviceOffering(params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, int, error) {
	var pciDevicePciDeviceOfferings []view.PciDevicePciDeviceOfferingRefInventoryView
	total, err := cli.Page("v1/pci-devices/pci-devices/pci-device-offerings", params, &pciDevicePciDeviceOfferings)
	return pciDevicePciDeviceOfferings, total, err
}
