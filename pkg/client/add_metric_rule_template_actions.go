// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddMetricRuleTemplate adds MetricRuleTemplate
func (cli *ZSClient) AddMetricRuleTemplate(params param.AddMetricRuleTemplateParam) (*view.AddMetricRuleTemplateEventView, error) {
	resp := view.AddMetricRuleTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/{monitorTemplateUuid}/metricrules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
