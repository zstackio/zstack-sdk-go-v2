// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBaremetalChassisState changes BaremetalChassisState
func (cli *ZSClient) ChangeBaremetalChassisState(uuid string, params param.ChangeBaremetalChassisStateParam) (*view.ChangeBaremetalChassisStateEventView, error) {
	resp := view.ChangeBaremetalChassisStateEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
