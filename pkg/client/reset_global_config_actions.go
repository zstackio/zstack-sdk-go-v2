// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResetGlobalConfig operates on ResetGlobalConfig
func (cli *ZSClient) ResetGlobalConfig(uuid string, params param.ResetGlobalConfigParam) (*view.ResetGlobalConfigEventView, error) {
	resp := view.ResetGlobalConfigEventView{}
	if err := cli.Put("v1/global-configurations/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
