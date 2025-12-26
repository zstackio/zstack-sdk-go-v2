// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySnmpAgent queries SnmpAgent list
func (cli *ZSClient) QuerySnmpAgent(params *param.QueryParam) ([]view.SnmpAgentInventoryView, error) {
	var resp []view.SnmpAgentInventoryView
	return resp, cli.List("v1/snmp/agent", params, &resp)
}
