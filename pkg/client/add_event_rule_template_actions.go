// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddEventRuleTemplate 操作AddEventRuleTemplate
func (cli *ZSClient) AddEventRuleTemplate(params param.AddEventRuleTemplateParam) (*view.AddEventRuleTemplateEventView, error) {
	resp := view.AddEventRuleTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/evenrules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

