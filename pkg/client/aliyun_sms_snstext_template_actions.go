// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunSmsSNSTextTemplate 更新AliyunSmsSNSTextTemplate
func (cli *ZSClient) UpdateAliyunSmsSNSTextTemplate(uuid string, params param.UpdateAliyunSmsSNSTextTemplateParam) (*view.UpdateAliyunSmsSNSTextTemplateEventView, error) {
	resp := view.UpdateAliyunSmsSNSTextTemplateEventView{}
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

