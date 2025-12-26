// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopSnmpAgent stops SnmpAgent
func (cli *ZSClient) StopSnmpAgent(uuid string, params param.StopSnmpAgentParam) (*view.StopSnmpAgentEventView, error) {
	resp := view.StopSnmpAgentEventView{}
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
