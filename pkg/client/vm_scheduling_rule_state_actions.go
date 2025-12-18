// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVmSchedulingRuleState 操作VmSchedulingRuleState
func (cli *ZSClient) ChangeVmSchedulingRuleState(uuid string, params param.ChangeVmSchedulingRuleStateParam) (*view.ChangeVmSchedulingRuleStateEventView, error) {
	resp := view.ChangeVmSchedulingRuleStateEventView{}
	if err := cli.Put("v1/vmSchedulingRule/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

