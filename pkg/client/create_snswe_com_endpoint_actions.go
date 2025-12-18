// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSWeComEndpoint creates SNSWeComEndpoint
func (cli *ZSClient) CreateSNSWeComEndpoint(params param.CreateSNSWeComEndpointParam) (*view.CreateSNSWeComEndpointEventView, error) {
	resp := view.CreateSNSWeComEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
