// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSApplicationEndpoint queries SNSApplicationEndpoint list
func (cli *ZSClient) QuerySNSApplicationEndpoint(params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, error) {
	var resp []view.SNSApplicationEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints", params, &resp)
}
