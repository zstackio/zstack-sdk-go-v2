// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSWeComEndpoint creates SNSWeComEndpoint
func (cli *ZSClient) CreateSNSWeComEndpoint(params param.CreateSNSWeComEndpointParam) (*view.CreateSNSWeComEndpointEventView, error) {
	resp := view.CreateSNSWeComEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
