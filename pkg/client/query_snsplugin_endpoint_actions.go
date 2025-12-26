// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSPluginEndpoint queries SNSPluginEndpoint list
func (cli *ZSClient) QuerySNSPluginEndpoint(params *param.QueryParam) ([]view.SNSPluginEndpointInventoryView, error) {
	var resp []view.SNSPluginEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/plugin", params, &resp)
}
