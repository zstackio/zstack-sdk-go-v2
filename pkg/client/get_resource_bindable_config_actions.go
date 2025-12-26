// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceBindableConfig gets ResourceBindableConfig by uuid
func (cli *ZSClient) GetResourceBindableConfig(uuid string) (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.Get("v1/resource-configurations/bindable", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
