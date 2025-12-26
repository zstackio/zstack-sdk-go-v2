// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostIommuState gets HostIommuState by uuid
func (cli *ZSClient) GetHostIommuState(uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
