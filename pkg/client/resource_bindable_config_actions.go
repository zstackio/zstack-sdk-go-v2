// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceBindableConfig 获取ResourceBindableConfig详情
func (cli *ZSClient) GetResourceBindableConfig(uuid string) (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.Get("v1/resource-configurations/bindable", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

