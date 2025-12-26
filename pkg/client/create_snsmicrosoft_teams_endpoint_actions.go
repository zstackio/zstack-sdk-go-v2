// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSMicrosoftTeamsEndpoint creates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) CreateSNSMicrosoftTeamsEndpoint(params param.CreateSNSMicrosoftTeamsEndpointParam) (*view.CreateSNSMicrosoftTeamsEndpointEventView, error) {
	resp := view.CreateSNSMicrosoftTeamsEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
