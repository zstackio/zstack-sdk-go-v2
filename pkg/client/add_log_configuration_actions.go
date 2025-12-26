// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddLogConfiguration adds LogConfiguration
func (cli *ZSClient) AddLogConfiguration(params param.AddLogConfigurationParam) (*view.AddLogConfigurationEventView, error) {
	resp := view.AddLogConfigurationEventView{}
	if err := cli.Post("v1/log/configurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
