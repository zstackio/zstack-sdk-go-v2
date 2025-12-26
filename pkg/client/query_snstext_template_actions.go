// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSTextTemplate queries SNSTextTemplate list
func (cli *ZSClient) QuerySNSTextTemplate(params *param.QueryParam) ([]view.SNSTextTemplateInventoryView, error) {
	var resp []view.SNSTextTemplateInventoryView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates", params, &resp)
}
