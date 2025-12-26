// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SNSDingTalkTestConnection operates on SNSDingTalkTestConnection
func (cli *ZSClient) SNSDingTalkTestConnection(params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
