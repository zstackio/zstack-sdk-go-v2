// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSSnmpTestConnection 操作SNSSnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

