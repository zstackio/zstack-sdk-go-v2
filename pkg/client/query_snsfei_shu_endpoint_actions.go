// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSFeiShuEndpoint queries SNSFeiShuEndpoint list
func (cli *ZSClient) QuerySNSFeiShuEndpoint(params *param.QueryParam) ([]view.SNSFeiShuEndpointInventoryView, error) {
	var resp []view.SNSFeiShuEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu", params, &resp)
}
