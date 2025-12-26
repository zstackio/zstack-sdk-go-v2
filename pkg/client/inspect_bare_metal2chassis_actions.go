// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// InspectBareMetal2Chassis operates on InspectBareMetal2Chassis
func (cli *ZSClient) InspectBareMetal2Chassis(uuid string, params param.InspectBareMetal2ChassisParam) (*view.InspectBareMetal2ChassisEventView, error) {
	resp := view.InspectBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
