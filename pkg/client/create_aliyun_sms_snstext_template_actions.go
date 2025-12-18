// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunSmsSNSTextTemplate creates AliyunSmsSNSTextTemplate
func (cli *ZSClient) CreateAliyunSmsSNSTextTemplate(params param.CreateAliyunSmsSNSTextTemplateParam) (*view.CreateSNSTextTemplateEventView, error) {
	resp := view.CreateSNSTextTemplateEventView{}
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates/aliyun-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
