// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleSchedulerJobTrigger creates AutoScalingRuleSchedulerJobTrigger
func (cli *ZSClient) CreateAutoScalingRuleSchedulerJobTrigger(params param.CreateAutoScalingRuleSchedulerJobTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	var resp view.CreateAutoScalingRuleTriggerEventView
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
