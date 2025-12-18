// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(params param.CreateOAuthClientParam) (*view.CreateOAuthClientEventView, error) {
	resp := view.CreateOAuthClientEventView{}
	if err := cli.Post("v1/create/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
