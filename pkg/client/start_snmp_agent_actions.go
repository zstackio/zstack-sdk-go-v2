// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartSnmpAgent starts SnmpAgent
func (cli *ZSClient) StartSnmpAgent(uuid string, params param.StartSnmpAgentParam) (*view.StartSnmpAgentEventView, error) {
	resp := view.StartSnmpAgentEventView{}
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
