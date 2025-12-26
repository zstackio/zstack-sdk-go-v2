// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BatchAddBareMetal2IpmiChassis operates on BatchAddBareMetal2IpmiChassis
func (cli *ZSClient) BatchAddBareMetal2IpmiChassis(params param.BatchAddBareMetal2IpmiChassisParam) (*view.BatchAddBareMetal2ChassisEventView, error) {
	resp := view.BatchAddBareMetal2ChassisEventView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
