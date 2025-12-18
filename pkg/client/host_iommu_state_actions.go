// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostIommuState 获取HostIommuState详情
func (cli *ZSClient) GetHostIommuState(uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

