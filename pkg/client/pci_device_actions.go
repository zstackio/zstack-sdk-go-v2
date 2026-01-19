// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(uuid string, params param.UpdatePciDeviceParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.Put("v1/pci-device/pci-devices", uuid, map[string]interface{}{
		"updatePciDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-devices", uuid, string(deleteMode))
}
// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List("v1/pci-device/pci-devices", params, &resp)
}

func (cli *ZSClient) GetPciDevice(uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.Get("v1/pci-device/pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDevice Pagination
func (cli *ZSClient) PagePciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, int, error) {
	var pciDevices []view.PciDeviceInventoryView
	total, err := cli.Page("v1/pci-device/pci-devices", params, &pciDevices)
	return pciDevices, total, err
}
