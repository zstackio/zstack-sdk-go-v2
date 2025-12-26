// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSUniversalSmsEndpoint queries SNSUniversalSmsEndpoint list
func (cli *ZSClient) QuerySNSUniversalSmsEndpoint(params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp []view.SNSUniversalSmsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/universal-sms", params, &resp)
}
