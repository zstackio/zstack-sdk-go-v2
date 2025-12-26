// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSHttpEndpoint creates SNSHttpEndpoint
func (cli *ZSClient) CreateSNSHttpEndpoint(params param.CreateSNSHttpEndpointParam) (*view.CreateSNSHttpEndpointEventView, error) {
	resp := view.CreateSNSHttpEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/http", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
