// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSHttpEndpoint queries SNSHttpEndpoint list
func (cli *ZSClient) QuerySNSHttpEndpoint(params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, error) {
	var resp []view.SNSHttpEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/http", params, &resp)
}
