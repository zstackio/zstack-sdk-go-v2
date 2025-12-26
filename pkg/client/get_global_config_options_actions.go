// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetGlobalConfigOptions gets GlobalConfigOptions by uuid
func (cli *ZSClient) GetGlobalConfigOptions(uuid string) (*view.GetGlobalConfigOptionsView, error) {
	var resp view.GetGlobalConfigOptionsView
	if err := cli.Get("v1/global-configurations/{category}/{name}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
