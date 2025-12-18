// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBareMetal2IpmiChassisHardwareInfo 创建BareMetal2IpmiChassisHardwareInfo
func (cli *ZSClient) CreateBareMetal2IpmiChassisHardwareInfo(params param.CreateBareMetal2IpmiChassisHardwareInfoParam) (*view.CreateBareMetal2ChassisHardwareView, error) {
	resp := view.CreateBareMetal2ChassisHardwareView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/hardwareinfos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

