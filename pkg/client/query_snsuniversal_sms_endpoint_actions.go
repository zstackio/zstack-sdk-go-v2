// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSUniversalSmsEndpoint queries SNSUniversalSmsEndpoint list
func (cli *ZSClient) QuerySNSUniversalSmsEndpoint(params param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp []view.SNSUniversalSmsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/universal-sms", &params, &resp)
}
