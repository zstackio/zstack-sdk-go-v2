// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// InspectBaremetalChassis operates on InspectBaremetalChassis
func (cli *ZSClient) InspectBaremetalChassis(uuid string, params param.InspectBaremetalChassisParam) (*view.InspectBaremetalChassisEventView, error) {
	resp := view.InspectBaremetalChassisEventView{}
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
