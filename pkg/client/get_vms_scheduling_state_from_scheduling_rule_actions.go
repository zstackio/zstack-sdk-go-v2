// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmsSchedulingStateFromSchedulingRule gets VmsSchedulingStateFromSchedulingRule by uuid
func (cli *ZSClient) GetVmsSchedulingStateFromSchedulingRule(uuid string) (*view.GetVmsSchedulingStateFromSchedulingRuleView, error) {
	var resp view.GetVmsSchedulingStateFromSchedulingRuleView
	if err := cli.Get("v1/get/vms/schedulingState/from/SchedulingRule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
