// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSEmailTestConnection operates on SNSEmailTestConnection
func (cli *ZSClient) SNSEmailTestConnection(params param.SNSEmailTestConnectionParam) (*view.SNSEmailTestConnectionEventView, error) {
	resp := view.SNSEmailTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/email/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
