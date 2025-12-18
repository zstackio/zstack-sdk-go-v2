// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSnmpAgent updates SnmpAgent
func (cli *ZSClient) UpdateSnmpAgent(uuid string, params param.UpdateSnmpAgentParam) (*view.UpdateSnmpAgentEventView, error) {
	resp := view.UpdateSnmpAgentEventView{}
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
