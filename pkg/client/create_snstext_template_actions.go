// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSTextTemplate creates SNSTextTemplate
func (cli *ZSClient) CreateSNSTextTemplate(params param.CreateSNSTextTemplateParam) (*view.CreateSNSTextTemplateEventView, error) {
	resp := view.CreateSNSTextTemplateEventView{}
	if err := cli.Post("v1/zwatch/alarms/sns/text-templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
