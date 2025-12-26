// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSDingTalkEndpoint creates SNSDingTalkEndpoint
func (cli *ZSClient) CreateSNSDingTalkEndpoint(params param.CreateSNSDingTalkEndpointParam) (*view.CreateSNSDingTalkEndpointEventView, error) {
	resp := view.CreateSNSDingTalkEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
