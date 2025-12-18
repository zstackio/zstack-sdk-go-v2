// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBareMetal2ChassisPowerStatus 获取BareMetal2ChassisPowerStatus详情
func (cli *ZSClient) GetBareMetal2ChassisPowerStatus(uuid string) (*view.GetBareMetal2ChassisPowerStatusView, error) {
	var resp view.GetBareMetal2ChassisPowerStatusView
	if err := cli.Get("v1/baremetal2/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

