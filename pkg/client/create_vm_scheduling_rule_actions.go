// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmSchedulingRule creates VmSchedulingRule
func (cli *ZSClient) CreateVmSchedulingRule(params param.CreateVmSchedulingRuleParam) (*view.CreateAffinityGroupEventView, error) {
	resp := view.CreateAffinityGroupEventView{}
	if err := cli.Post("v1/vmsSchedulingRule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
