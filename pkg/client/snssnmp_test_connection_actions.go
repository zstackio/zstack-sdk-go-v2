// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSSnmpTestConnection operates on SNSSnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
