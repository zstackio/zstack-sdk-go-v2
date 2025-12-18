// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSMicrosoftTeamsTestConnection 操作SNSMicrosoftTeamsTestConnection
func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

