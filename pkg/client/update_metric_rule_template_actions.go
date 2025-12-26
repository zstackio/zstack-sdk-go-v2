// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateMetricRuleTemplate updates MetricRuleTemplate
func (cli *ZSClient) UpdateMetricRuleTemplate(uuid string, params param.UpdateMetricRuleTemplateParam) (*view.UpdateMetricRuleTemplateEventView, error) {
	resp := view.UpdateMetricRuleTemplateEventView{}
	if err := cli.Put("v1/zwatch/monitortemplates/metricrules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
