// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunSmsSNSTextTemplate updates AliyunSmsSNSTextTemplate
func (cli *ZSClient) UpdateAliyunSmsSNSTextTemplate(uuid string, params param.UpdateAliyunSmsSNSTextTemplateParam) (*view.UpdateAliyunSmsSNSTextTemplateEventView, error) {
	resp := view.UpdateAliyunSmsSNSTextTemplateEventView{}
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
