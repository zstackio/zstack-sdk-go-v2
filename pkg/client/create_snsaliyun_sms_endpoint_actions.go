// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSAliyunSmsEndpoint creates SNSAliyunSmsEndpoint
func (cli *ZSClient) CreateSNSAliyunSmsEndpoint(params param.CreateSNSAliyunSmsEndpointParam) (*view.CreateSNSAliyunSmsEndpointEventView, error) {
	resp := view.CreateSNSAliyunSmsEndpointEventView{}
	if err := cli.Post("v1/sns/sms-endpoints/aliyunsms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
