// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(params *param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp []view.SNSAliyunSmsEndpointInventoryView
	return resp, cli.List("v1/sns/sms-endpoints", params, &resp)
}
