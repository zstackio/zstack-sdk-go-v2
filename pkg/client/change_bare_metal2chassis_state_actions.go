// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBareMetal2ChassisState changes BareMetal2ChassisState
func (cli *ZSClient) ChangeBareMetal2ChassisState(uuid string, params param.ChangeBareMetal2ChassisStateParam) (*view.ChangeBareMetal2ChassisStateEventView, error) {
	resp := view.ChangeBareMetal2ChassisStateEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
