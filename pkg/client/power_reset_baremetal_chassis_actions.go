// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerResetBaremetalChassis operates on PowerResetBaremetalChassis
func (cli *ZSClient) PowerResetBaremetalChassis(uuid string, params param.PowerResetBaremetalChassisParam) (*view.PowerResetBaremetalChassisEventView, error) {
	resp := view.PowerResetBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
