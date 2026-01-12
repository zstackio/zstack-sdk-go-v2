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

func (cli *ZSClient) GetPciDeviceSpec(uuid string) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.Get("v1/pci-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(uuid string, params param.UpdatePciDeviceSpecParam) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.UpdatePciDeviceSpecEventView
	err := cli.PutWithSpec("v1/pci-device-specs", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
