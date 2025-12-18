// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerResetBareMetal2Chassis operates on PowerResetBareMetal2Chassis
func (cli *ZSClient) PowerResetBareMetal2Chassis(uuid string, params param.PowerResetBareMetal2ChassisParam) (*view.PowerResetBareMetal2ChassisEventView, error) {
	resp := view.PowerResetBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
