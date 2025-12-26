// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSPluginEndpoint creates SNSPluginEndpoint
func (cli *ZSClient) CreateSNSPluginEndpoint(params param.CreateSNSPluginEndpointParam) (*view.CreateSNSPluginEndpointEventView, error) {
	resp := view.CreateSNSPluginEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
