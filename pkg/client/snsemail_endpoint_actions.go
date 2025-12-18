// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSEmailEndpoint 创建SNSEmailEndpoint
func (cli *ZSClient) CreateSNSEmailEndpoint(params param.CreateSNSEmailEndpointParam) (*view.CreateSNSApplicationEndpointEventView, error) {
	resp := view.CreateSNSApplicationEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

