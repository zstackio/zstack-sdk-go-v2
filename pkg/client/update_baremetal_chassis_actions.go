// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBaremetalChassis updates BaremetalChassis
func (cli *ZSClient) UpdateBaremetalChassis(uuid string, params param.UpdateBaremetalChassisParam) (*view.UpdateBaremetalChassisEventView, error) {
	resp := view.UpdateBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
