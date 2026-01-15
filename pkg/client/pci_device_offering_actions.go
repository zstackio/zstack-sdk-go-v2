// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePciDeviceOffering deletes PciDeviceOffering
func (cli *ZSClient) DeletePciDeviceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-device-offerings", uuid, string(deleteMode))
}
// CreatePciDeviceOffering creates PciDeviceOffering
func (cli *ZSClient) CreatePciDeviceOffering(params param.CreatePciDeviceOfferingParam) (*view.PciDeviceOfferingInventoryView, error) {
	resp := view.PciDeviceOfferingInventoryView{}
	if err := cli.Post("v1/pci-device/pci-device-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPciDeviceOffering queries PciDeviceOffering list
func (cli *ZSClient) QueryPciDeviceOffering(params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, error) {
	var resp []view.PciDeviceOfferingInventoryView
	return resp, cli.List("v1/pci-device/pci-device-offerings", params, &resp)
}

// PagePciDeviceOffering Pagination
func (cli *ZSClient) PagePciDeviceOffering(params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, int, error) {
	var pciDeviceOfferings []view.PciDeviceOfferingInventoryView
	total, err := cli.Page("v1/pci-device/pci-device-offerings", params, &pciDeviceOfferings)
	return pciDeviceOfferings, total, err
}
