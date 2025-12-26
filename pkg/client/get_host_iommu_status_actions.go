// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostIommuStatus gets HostIommuStatus by uuid
func (cli *ZSClient) GetHostIommuStatus(uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
