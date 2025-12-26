// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePluginSecretResourcePool creates PluginSecretResourcePool
func (cli *ZSClient) CreatePluginSecretResourcePool(params param.CreatePluginSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
