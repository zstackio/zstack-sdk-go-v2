// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleSchedulerJobTrigger creates AutoScalingRuleSchedulerJobTrigger
func (cli *ZSClient) CreateAutoScalingRuleSchedulerJobTrigger(ctx context.Context, schedulerJobUuid string, ruleUuid string, params param.CreateAutoScalingRuleSchedulerJobTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	resp := view.AutoScalingRuleTriggerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobs/%s/autoscaling/rules/%s", schedulerJobUuid, ruleUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
