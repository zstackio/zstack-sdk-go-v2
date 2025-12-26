// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSHttpTestConnection operates on SNSHttpTestConnection
func (cli *ZSClient) SNSHttpTestConnection(params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/http/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
