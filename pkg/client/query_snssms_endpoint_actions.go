// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(params param.QueryParam) ([]view.SNSAliyunSmsEndpointInventoryView, error) {
	var resp []view.SNSAliyunSmsEndpointInventoryView
	return resp, cli.List("v1/sns/sms-endpoints", &params, &resp)
}
