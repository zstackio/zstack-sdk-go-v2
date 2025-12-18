// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSHttpEndpoint queries SNSHttpEndpoint list
func (cli *ZSClient) QuerySNSHttpEndpoint(params param.QueryParam) ([]view.SNSHttpEndpointInventoryView, error) {
	var resp []view.SNSHttpEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/http", &params, &resp)
}
