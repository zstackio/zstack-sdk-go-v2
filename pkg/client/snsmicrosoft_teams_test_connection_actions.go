// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSMicrosoftTeamsTestConnection operates on SNSMicrosoftTeamsTestConnection
func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
