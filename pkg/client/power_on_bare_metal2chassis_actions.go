// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerOnBareMetal2Chassis operates on PowerOnBareMetal2Chassis
func (cli *ZSClient) PowerOnBareMetal2Chassis(uuid string, params param.PowerOnBareMetal2ChassisParam) (*view.PowerOnBareMetal2ChassisEventView, error) {
	resp := view.PowerOnBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
