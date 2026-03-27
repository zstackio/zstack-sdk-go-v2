// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAutoScalingRuleAlarmTrigger creates AutoScalingRuleAlarmTrigger
func (cli *ZSClient) CreateAutoScalingRuleAlarmTrigger(ctx context.Context, params param.CreateAutoScalingRuleAlarmTriggerParam) (*view.AutoScalingRuleTriggerInventoryView, error) {
	resp := view.AutoScalingRuleTriggerInventoryView{}
	if err := cli.Post(ctx, "v1/zwatch/alarms/{alarmUuid}/autoscaling/rules/{ruleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
