// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSWeComTestConnection operates on SNSWeComTestConnection
func (cli *ZSClient) SNSWeComTestConnection(params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
