// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(uuid string, params param.UpdatePciDeviceParam) (*view.PciDeviceInventoryView, error) {
	var resp view.UpdatePciDeviceEventView
	if err := cli.Put("v1/pci-device/pci-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-devices/{uuid}", uuid, string(deleteMode))
}
// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List("v1/pci-device/pci-devices", params, &resp)
}
