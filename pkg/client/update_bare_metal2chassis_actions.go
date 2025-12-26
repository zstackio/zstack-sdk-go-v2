// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBareMetal2Chassis updates BareMetal2Chassis
func (cli *ZSClient) UpdateBareMetal2Chassis(uuid string, params param.UpdateBareMetal2ChassisParam) (*view.UpdateBareMetal2ChassisEventView, error) {
	resp := view.UpdateBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
