// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerOffBareMetal2Chassis operates on PowerOffBareMetal2Chassis
func (cli *ZSClient) PowerOffBareMetal2Chassis(uuid string, params param.PowerOffBareMetal2ChassisParam) (*view.PowerOffBareMetal2ChassisEventView, error) {
	resp := view.PowerOffBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
