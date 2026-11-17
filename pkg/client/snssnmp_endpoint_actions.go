// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SNSSnmpTestConnection operates on SnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSNSSnmpEndpoint creates SNSSnmpEndpoint
func (cli *ZSClient) CreateSNSSnmpEndpoint() (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
