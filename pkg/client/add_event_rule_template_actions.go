// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddEventRuleTemplate adds EventRuleTemplate
func (cli *ZSClient) AddEventRuleTemplate(params param.AddEventRuleTemplateParam) (*view.AddEventRuleTemplateEventView, error) {
	resp := view.AddEventRuleTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/evenrules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
