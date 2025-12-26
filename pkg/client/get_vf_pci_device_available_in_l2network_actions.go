// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVfPciDeviceAvailableInL2Network gets VfPciDeviceAvailableInL2Network by uuid
func (cli *ZSClient) GetVfPciDeviceAvailableInL2Network(uuid string) (*view.GetVfPciDeviceAvailableInL2NetworkView, error) {
	var resp view.GetVfPciDeviceAvailableInL2NetworkView
	if err := cli.Get("v1/l2-networks/vf-pci-devices-available", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
