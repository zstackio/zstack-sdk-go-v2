// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEventRuleTemplate queries EventRuleTemplate list
func (cli *ZSClient) QueryEventRuleTemplate(params *param.QueryParam) ([]view.EventRuleTemplateInventoryView, error) {
	var resp []view.EventRuleTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates/evenrules", params, &resp)
}
