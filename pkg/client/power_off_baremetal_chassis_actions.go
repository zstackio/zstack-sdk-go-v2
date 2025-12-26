// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerOffBaremetalChassis operates on PowerOffBaremetalChassis
func (cli *ZSClient) PowerOffBaremetalChassis(uuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
