// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UngenerateSriovPciDevices 操作UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(uuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.Put("v1/pci-devices/{pciDeviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

