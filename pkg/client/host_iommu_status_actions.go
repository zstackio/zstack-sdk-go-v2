// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostIommuStatus 获取HostIommuStatus详情
func (cli *ZSClient) GetHostIommuStatus(uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.Get("v1/pci-device/hosts/{uuid}/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

