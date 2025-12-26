// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSFeiShuTestConnection operates on SNSFeiShuTestConnection
func (cli *ZSClient) SNSFeiShuTestConnection(params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
