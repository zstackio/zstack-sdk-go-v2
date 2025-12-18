// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSDingTalkTestConnection 操作SNSDingTalkTestConnection
func (cli *ZSClient) SNSDingTalkTestConnection(params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

