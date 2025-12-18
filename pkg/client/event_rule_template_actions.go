// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventRuleTemplate 查询EventRuleTemplate列表
func (cli *ZSClient) QueryEventRuleTemplate(params param.QueryParam) ([]view.QueryEventRuleTemplateView, error) {
	var resp []view.QueryEventRuleTemplateView
	return resp, cli.List("v1/zwatch/monitortemplates/evenrules", &params, &resp)
}

