// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncAINginxConfiguration operates on SyncAINginxConfiguration
func (cli *ZSClient) SyncAINginxConfiguration(params param.SyncAINginxConfigurationParam) (*view.SyncAINginxConfigurationView, error) {
	resp := view.SyncAINginxConfigurationView{}
	if err := cli.Post("v1/ai/nginx/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
