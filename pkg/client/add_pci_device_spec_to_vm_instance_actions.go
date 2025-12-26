// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddPciDeviceSpecToVmInstance adds PciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(params param.AddPciDeviceSpecToVmInstanceParam) (*view.AddPciDeviceSpecToVmInstanceEventView, error) {
	resp := view.AddPciDeviceSpecToVmInstanceEventView{}
	if err := cli.Post("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
