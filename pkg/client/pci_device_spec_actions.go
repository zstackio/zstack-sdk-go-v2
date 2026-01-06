// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPciDeviceSpec queries PciDeviceSpec list
func (cli *ZSClient) QueryPciDeviceSpec(params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, error) {
	var resp []view.PciDeviceSpecInventoryView
	return resp, cli.List("v1/pci-device-specs", params, &resp)
}
// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(uuid string, params param.UpdatePciDeviceSpecParam) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.UpdatePciDeviceSpecEventView
	if err := cli.Put("v1/pci-device-specs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
