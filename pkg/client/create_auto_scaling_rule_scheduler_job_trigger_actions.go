// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAutoScalingRuleSchedulerJobTrigger creates AutoScalingRuleSchedulerJobTrigger
func (cli *ZSClient) CreateAutoScalingRuleSchedulerJobTrigger(params param.CreateAutoScalingRuleSchedulerJobTriggerParam) (*view.CreateAutoScalingRuleTriggerEventView, error) {
	resp := view.CreateAutoScalingRuleTriggerEventView{}
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
