// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMetricRuleTemplate queries MetricRuleTemplate list
func (cli *ZSClient) QueryMetricRuleTemplate(params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, error) {
	var resp []view.MetricRuleTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates/metricrules", params, &resp)
}
