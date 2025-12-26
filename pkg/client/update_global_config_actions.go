// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateGlobalConfig updates GlobalConfig
func (cli *ZSClient) UpdateGlobalConfig(uuid string, params param.UpdateGlobalConfigParam) (*view.UpdateGlobalConfigEventView, error) {
	resp := view.UpdateGlobalConfigEventView{}
	if err := cli.Put("v1/global-configurations/{category}/{name}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
