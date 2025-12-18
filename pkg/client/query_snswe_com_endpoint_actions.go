// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSWeComEndpoint queries SNSWeComEndpoint list
func (cli *ZSClient) QuerySNSWeComEndpoint(params param.QueryParam) ([]view.SNSWeComEndpointInventoryView, error) {
	var resp []view.SNSWeComEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com", &params, &resp)
}
