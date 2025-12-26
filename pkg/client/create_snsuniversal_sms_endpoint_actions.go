// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSUniversalSmsEndpoint creates SNSUniversalSmsEndpoint
func (cli *ZSClient) CreateSNSUniversalSmsEndpoint(params param.CreateSNSUniversalSmsEndpointParam) (*view.CreateSNSUniversalSmsEndpointEventView, error) {
	resp := view.CreateSNSUniversalSmsEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/universal-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
