// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddBareMetal2IpmiChassis adds BareMetal2IpmiChassis
func (cli *ZSClient) AddBareMetal2IpmiChassis(params param.AddBareMetal2IpmiChassisParam) (*view.AddBareMetal2ChassisEventView, error) {
	resp := view.AddBareMetal2ChassisEventView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
