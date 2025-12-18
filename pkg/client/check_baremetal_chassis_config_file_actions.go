// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckBaremetalChassisConfigFile 操作CheckBaremetalChassisConfigFile
func (cli *ZSClient) CheckBaremetalChassisConfigFile(params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.Post("v1/baremetal/chassis/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

