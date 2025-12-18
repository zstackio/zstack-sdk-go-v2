// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLogConfiguration 操作AddLogConfiguration
func (cli *ZSClient) AddLogConfiguration(params param.AddLogConfigurationParam) (*view.AddLogConfigurationEventView, error) {
	resp := view.AddLogConfigurationEventView{}
	if err := cli.Post("v1/log/configurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

