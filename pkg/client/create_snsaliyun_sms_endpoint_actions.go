// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSAliyunSmsEndpoint creates SNSAliyunSmsEndpoint
func (cli *ZSClient) CreateSNSAliyunSmsEndpoint(params param.CreateSNSAliyunSmsEndpointParam) (*view.CreateSNSAliyunSmsEndpointEventView, error) {
	resp := view.CreateSNSAliyunSmsEndpointEventView{}
	if err := cli.Post("v1/sns/sms-endpoints/aliyunsms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
