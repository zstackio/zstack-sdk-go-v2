// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSMicrosoftTeamsEndpoint updates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) UpdateSNSMicrosoftTeamsEndpoint(uuid string, params param.UpdateSNSMicrosoftTeamsEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/microsoft-teams/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
