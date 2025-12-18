// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SNSHttpTestConnection 操作SNSHttpTestConnection
func (cli *ZSClient) SNSHttpTestConnection(params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.Post("v1/sns/application-endpoints/http/test-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

