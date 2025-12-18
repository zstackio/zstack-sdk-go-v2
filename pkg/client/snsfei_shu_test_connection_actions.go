// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSFeiShuTestConnection 操作SNSFeiShuTestConnection
func (cli *ZSClient) SNSFeiShuTestConnection(params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

