// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSTextTemplate updates SNSTextTemplate
func (cli *ZSClient) UpdateSNSTextTemplate(uuid string, params param.UpdateSNSTextTemplateParam) (*view.UpdateSNSTextTemplateEventView, error) {
	resp := view.UpdateSNSTextTemplateEventView{}
	if err := cli.Put("v1/zwatch/alarms/sns/text-templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
