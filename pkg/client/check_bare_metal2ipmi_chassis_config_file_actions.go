// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckBareMetal2IpmiChassisConfigFile 操作CheckBareMetal2IpmiChassisConfigFile
func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile(params param.CheckBareMetal2IpmiChassisConfigFileParam) (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

