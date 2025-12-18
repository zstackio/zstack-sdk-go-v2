// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSWeComTestConnection 操作SNSWeComTestConnection
func (cli *ZSClient) SNSWeComTestConnection(params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

