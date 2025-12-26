// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSFeiShuEndpoint creates SNSFeiShuEndpoint
func (cli *ZSClient) CreateSNSFeiShuEndpoint(params param.CreateSNSFeiShuEndpointParam) (*view.CreateSNSFeiShuEndpointEventView, error) {
	resp := view.CreateSNSFeiShuEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
