// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StopSnmpAgent 停止SnmpAgent
func (cli *ZSClient) StopSnmpAgent(uuid string, params param.StopSnmpAgentParam) (*view.StopSnmpAgentEventView, error) {
	resp := view.StopSnmpAgentEventView{}
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

