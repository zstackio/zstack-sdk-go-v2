// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSMicrosoftTeamsEndpoint queries SNSMicrosoftTeamsEndpoint list
func (cli *ZSClient) QuerySNSMicrosoftTeamsEndpoint(params *param.QueryParam) ([]view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp []view.SNSMicrosoftTeamsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/microsoft-teams", params, &resp)
}
