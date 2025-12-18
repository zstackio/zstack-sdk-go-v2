// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSPluginEndpoint creates SNSPluginEndpoint
func (cli *ZSClient) CreateSNSPluginEndpoint(params param.CreateSNSPluginEndpointParam) (*view.CreateSNSPluginEndpointEventView, error) {
	resp := view.CreateSNSPluginEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
