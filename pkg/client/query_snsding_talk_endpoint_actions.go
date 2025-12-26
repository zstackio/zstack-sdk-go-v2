// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSDingTalkEndpoint queries SNSDingTalkEndpoint list
func (cli *ZSClient) QuerySNSDingTalkEndpoint(params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, error) {
	var resp []view.SNSDingTalkEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk", params, &resp)
}
