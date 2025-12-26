// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BatchCreateBaremetalChassis operates on BatchCreateBaremetalChassis
func (cli *ZSClient) BatchCreateBaremetalChassis(params param.BatchCreateBaremetalChassisParam) (*view.BatchCreateBaremetalChassisEventView, error) {
	resp := view.BatchCreateBaremetalChassisEventView{}
	if err := cli.Post("v1/baremetal/chassis/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
