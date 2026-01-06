// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePciDeviceOffering deletes PciDeviceOffering
func (cli *ZSClient) DeletePciDeviceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-device-offerings/{uuid}", uuid, string(deleteMode))
}
// CreatePciDeviceOffering creates PciDeviceOffering
func (cli *ZSClient) CreatePciDeviceOffering(params param.CreatePciDeviceOfferingParam) (*view.PciDeviceOfferingInventoryView, error) {
	var resp view.CreatePciDeviceOfferingEventView
	if err := cli.Post("v1/pci-device/pci-device-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryPciDeviceOffering queries PciDeviceOffering list
func (cli *ZSClient) QueryPciDeviceOffering(params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, error) {
	var resp []view.PciDeviceOfferingInventoryView
	return resp, cli.List("v1/pci-device/pci-device-offerings", params, &resp)
}
