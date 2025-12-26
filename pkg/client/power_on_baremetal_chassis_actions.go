// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerOnBaremetalChassis operates on PowerOnBaremetalChassis
func (cli *ZSClient) PowerOnBaremetalChassis(uuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
