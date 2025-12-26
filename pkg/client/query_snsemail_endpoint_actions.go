// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSEmailEndpoint queries SNSEmailEndpoint list
func (cli *ZSClient) QuerySNSEmailEndpoint(params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, error) {
	var resp []view.SNSEmailEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails", params, &resp)
}
