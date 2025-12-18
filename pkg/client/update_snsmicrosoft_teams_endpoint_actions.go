// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSNSMicrosoftTeamsEndpoint updates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) UpdateSNSMicrosoftTeamsEndpoint(uuid string, params param.UpdateSNSMicrosoftTeamsEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/microsoft-teams/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
