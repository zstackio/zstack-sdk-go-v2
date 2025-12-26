// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckBaremetalChassisConfigFile operates on CheckBaremetalChassisConfigFile
func (cli *ZSClient) CheckBaremetalChassisConfigFile(params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.Post("v1/baremetal/chassis/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
