// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BatchAddBareMetal2IpmiChassis operates on BatchAddBareMetal2IpmiChassis
func (cli *ZSClient) BatchAddBareMetal2IpmiChassis(params param.BatchAddBareMetal2IpmiChassisParam) (*view.BatchAddBareMetal2ChassisEventView, error) {
	resp := view.BatchAddBareMetal2ChassisEventView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
