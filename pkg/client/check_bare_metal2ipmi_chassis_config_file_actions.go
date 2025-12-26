// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckBareMetal2IpmiChassisConfigFile operates on CheckBareMetal2IpmiChassisConfigFile
func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile(params param.CheckBareMetal2IpmiChassisConfigFileParam) (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
