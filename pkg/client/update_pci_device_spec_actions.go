// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(uuid string, params param.UpdatePciDeviceSpecParam) (*view.UpdatePciDeviceSpecEventView, error) {
	resp := view.UpdatePciDeviceSpecEventView{}
	if err := cli.Put("v1/pci-device-specs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
