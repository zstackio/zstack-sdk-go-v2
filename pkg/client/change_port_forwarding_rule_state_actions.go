// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangePortForwardingRuleState changes PortForwardingRuleState
func (cli *ZSClient) ChangePortForwardingRuleState(uuid string, params param.ChangePortForwardingRuleStateParam) (*view.ChangePortForwardingRuleStateEventView, error) {
	resp := view.ChangePortForwardingRuleStateEventView{}
	if err := cli.Put("v1/port-forwarding/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
