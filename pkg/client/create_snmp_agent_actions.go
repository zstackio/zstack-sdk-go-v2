// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSnmpAgent creates SnmpAgent
func (cli *ZSClient) CreateSnmpAgent(params param.CreateSnmpAgentParam) (*view.CreateSnmpAgentEventView, error) {
	resp := view.CreateSnmpAgentEventView{}
	if err := cli.Post("v1/snmp/agent", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
