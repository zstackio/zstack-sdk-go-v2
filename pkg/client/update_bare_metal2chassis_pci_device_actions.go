// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBareMetal2ChassisPciDevice updates BareMetal2ChassisPciDevice
func (cli *ZSClient) UpdateBareMetal2ChassisPciDevice(uuid string, params param.UpdateBareMetal2ChassisPciDeviceParam) (*view.UpdateBareMetal2ChassisPciDeviceEventView, error) {
	resp := view.UpdateBareMetal2ChassisPciDeviceEventView{}
	if err := cli.Put("v1/baremetal2/chassis/pci-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
