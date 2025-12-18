// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSMicrosoftTeamsEndpoint creates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) CreateSNSMicrosoftTeamsEndpoint(params param.CreateSNSMicrosoftTeamsEndpointParam) (*view.CreateSNSMicrosoftTeamsEndpointEventView, error) {
	resp := view.CreateSNSMicrosoftTeamsEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
