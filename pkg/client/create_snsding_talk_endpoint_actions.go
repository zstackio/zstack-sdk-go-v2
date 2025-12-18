// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSDingTalkEndpoint creates SNSDingTalkEndpoint
func (cli *ZSClient) CreateSNSDingTalkEndpoint(params param.CreateSNSDingTalkEndpointParam) (*view.CreateSNSDingTalkEndpointEventView, error) {
	resp := view.CreateSNSDingTalkEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/ding-talk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
