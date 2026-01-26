// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleSchedulerJobTrigger creates AutoScalingRuleSchedulerJobTrigger
func (cli *ZSClient) CreateAutoScalingRuleSchedulerJobTrigger(params param.CreateAutoScalingRuleSchedulerJobTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	resp := view.AutoScalingRuleTriggerInventoryView{}
	if err := cli.Post("v1/scheduler/jobs/{schedulerJobUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
