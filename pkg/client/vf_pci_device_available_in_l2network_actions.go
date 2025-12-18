// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVfPciDeviceAvailableInL2Network 获取VfPciDeviceAvailableInL2Network详情
func (cli *ZSClient) GetVfPciDeviceAvailableInL2Network(uuid string) (*view.GetVfPciDeviceAvailableInL2NetworkView, error) {
	var resp view.GetVfPciDeviceAvailableInL2NetworkView
	if err := cli.Get("v1/l2-networks/vf-pci-devices-available", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

