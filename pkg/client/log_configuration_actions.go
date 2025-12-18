// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateLogConfiguration 更新LogConfiguration
func (cli *ZSClient) UpdateLogConfiguration(uuid string, params param.UpdateLogConfigurationParam) (*view.UpdateLogConfigurationEventView, error) {
	resp := view.UpdateLogConfigurationEventView{}
	if err := cli.Put("v1/log/configurations", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

