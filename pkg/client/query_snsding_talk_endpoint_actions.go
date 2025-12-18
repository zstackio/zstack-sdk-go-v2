// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSDingTalkEndpoint queries SNSDingTalkEndpoint list
func (cli *ZSClient) QuerySNSDingTalkEndpoint(params param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, error) {
	var resp []view.SNSDingTalkEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk", &params, &resp)
}
