// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEthernetVF queries EthernetVF list
func (cli *ZSClient) QueryEthernetVF(params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, error) {
	var resp []view.EthernetVfPciDeviceInventoryView
	return resp, cli.List("v1/pci-device/ethernet-vfs", params, &resp)
}
