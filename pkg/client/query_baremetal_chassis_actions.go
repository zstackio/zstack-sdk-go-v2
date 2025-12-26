// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBaremetalChassis queries BaremetalChassis list
func (cli *ZSClient) QueryBaremetalChassis(params *param.QueryParam) ([]view.BaremetalChassisInventoryView, error) {
	var resp []view.BaremetalChassisInventoryView
	return resp, cli.List("v1/baremetal/chassis", params, &resp)
}
