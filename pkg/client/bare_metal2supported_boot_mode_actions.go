// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBareMetal2SupportedBootMode 获取BareMetal2SupportedBootMode详情
func (cli *ZSClient) GetBareMetal2SupportedBootMode(uuid string) (*view.GetBareMetal2SupportedBootModeView, error) {
	var resp view.GetBareMetal2SupportedBootModeView
	if err := cli.Get("v1/baremetal2/chassis/supported-boot-modes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

