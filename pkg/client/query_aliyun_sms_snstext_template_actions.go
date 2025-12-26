// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunSmsSNSTextTemplate queries AliyunSmsSNSTextTemplate list
func (cli *ZSClient) QueryAliyunSmsSNSTextTemplate(params *param.QueryParam) ([]view.AliyunSmsSNSTextTemplateInventoryView, error) {
	var resp []view.AliyunSmsSNSTextTemplateInventoryView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp)
}
