// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSWeComEndpoint queries SNSWeComEndpoint list
func (cli *ZSClient) QuerySNSWeComEndpoint(params *param.QueryParam) ([]view.SNSWeComEndpointInventoryView, error) {
	var resp []view.SNSWeComEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com", params, &resp)
}
