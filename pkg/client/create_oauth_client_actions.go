// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(params param.CreateOAuthClientParam) (*view.CreateOAuthClientEventView, error) {
	resp := view.CreateOAuthClientEventView{}
	if err := cli.Post("v1/create/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
