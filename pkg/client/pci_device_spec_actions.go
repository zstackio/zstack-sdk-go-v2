// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPciDeviceSpec queries PciDeviceSpec list
func (cli *ZSClient) QueryPciDeviceSpec(params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, error) {
	var resp []view.PciDeviceSpecInventoryView
	return resp, cli.List("v1/pci-device-specs", params, &resp)
}

// PagePciDeviceSpec Pagination
func (cli *ZSClient) PagePciDeviceSpec(params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, int, error) {
	var pciDeviceSpecs []view.PciDeviceSpecInventoryView
	total, err := cli.Page("v1/pci-device-specs", params, &pciDeviceSpecs)
	return pciDeviceSpecs, total, err
}
// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(uuid string, params param.UpdatePciDeviceSpecParam) (*view.PciDeviceSpecInventoryView, error) {
	resp := view.PciDeviceSpecInventoryView{}
	if err := cli.Put("v1/pci-device-specs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
