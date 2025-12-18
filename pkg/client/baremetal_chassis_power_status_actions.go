// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBaremetalChassisPowerStatus 获取BaremetalChassisPowerStatus详情
func (cli *ZSClient) GetBaremetalChassisPowerStatus(uuid string) (*view.GetBaremetalChassisPowerStatusView, error) {
	var resp view.GetBaremetalChassisPowerStatusView
	if err := cli.Get("v1/baremetal/chassis/{uuid}/powerstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

