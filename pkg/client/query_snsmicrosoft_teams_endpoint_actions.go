// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSMicrosoftTeamsEndpoint queries SNSMicrosoftTeamsEndpoint list
func (cli *ZSClient) QuerySNSMicrosoftTeamsEndpoint(params param.QueryParam) ([]view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp []view.SNSMicrosoftTeamsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/microsoft-teams", &params, &resp)
}
