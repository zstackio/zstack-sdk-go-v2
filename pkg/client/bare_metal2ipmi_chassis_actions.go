// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBareMetal2IpmiChassis 更新BareMetal2IpmiChassis
func (cli *ZSClient) UpdateBareMetal2IpmiChassis(uuid string, params param.UpdateBareMetal2IpmiChassisParam) (*view.UpdateBareMetal2ChassisEventView, error) {
	resp := view.UpdateBareMetal2ChassisEventView{}
	if err := cli.Put("v1/baremetal2/chassis/ipmi/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

