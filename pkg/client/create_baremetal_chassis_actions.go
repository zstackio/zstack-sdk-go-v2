// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBaremetalChassis creates BaremetalChassis
func (cli *ZSClient) CreateBaremetalChassis(params param.CreateBaremetalChassisParam) (*view.CreateBaremetalChassisEventView, error) {
	resp := view.CreateBaremetalChassisEventView{}
	if err := cli.Post("v1/baremetal/chassis", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
