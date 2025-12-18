// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddBareMetal2IpmiChassis 操作AddBareMetal2IpmiChassis
func (cli *ZSClient) AddBareMetal2IpmiChassis(params param.AddBareMetal2IpmiChassisParam) (*view.AddBareMetal2ChassisEventView, error) {
	resp := view.AddBareMetal2ChassisEventView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

