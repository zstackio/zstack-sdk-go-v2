// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostPowerStatus 获取HostPowerStatus详情
func (cli *ZSClient) GetHostPowerStatus(uuid string) (*view.GetHostPowerStatusEventView, error) {
	var resp view.GetHostPowerStatusEventView
	if err := cli.Get("v1/hosts/power/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

