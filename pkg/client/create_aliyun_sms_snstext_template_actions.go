// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunSmsSNSTextTemplate creates AliyunSmsSNSTextTemplate
func (cli *ZSClient) CreateAliyunSmsSNSTextTemplate(params param.CreateAliyunSmsSNSTextTemplateParam) (*view.CreateSNSTextTemplateEventView, error) {
	resp := view.CreateSNSTextTemplateEventView{}
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
