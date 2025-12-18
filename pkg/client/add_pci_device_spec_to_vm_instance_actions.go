// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddPciDeviceSpecToVmInstance 操作AddPciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(params param.AddPciDeviceSpecToVmInstanceParam) (*view.AddPciDeviceSpecToVmInstanceEventView, error) {
	resp := view.AddPciDeviceSpecToVmInstanceEventView{}
	if err := cli.Post("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

