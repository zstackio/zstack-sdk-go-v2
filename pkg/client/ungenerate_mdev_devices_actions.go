// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UngenerateMdevDevices 操作UngenerateMdevDevices
func (cli *ZSClient) UngenerateMdevDevices(uuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

