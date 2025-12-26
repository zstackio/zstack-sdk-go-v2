// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEventRuleTemplate updates EventRuleTemplate
func (cli *ZSClient) UpdateEventRuleTemplate(uuid string, params param.UpdateEventRuleTemplateParam) (*view.UpdateEventRuleTemplateEventView, error) {
	resp := view.UpdateEventRuleTemplateEventView{}
	if err := cli.Put("v1/zwatch/monitortemplates/evenrules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
