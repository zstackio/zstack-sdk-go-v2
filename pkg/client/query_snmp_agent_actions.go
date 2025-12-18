// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySnmpAgent queries SnmpAgent list
func (cli *ZSClient) QuerySnmpAgent(params param.QueryParam) ([]view.SnmpAgentInventoryView, error) {
	var resp []view.SnmpAgentInventoryView
	return resp, cli.List("v1/snmp/agent", &params, &resp)
}
