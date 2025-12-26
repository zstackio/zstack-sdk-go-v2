// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmInstancePciDeviceSpecRef queries VmInstancePciDeviceSpecRef list
func (cli *ZSClient) QueryVmInstancePciDeviceSpecRef(params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstancePciDeviceSpecRefInventoryView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/pci-device-specs", params, &resp)
}
