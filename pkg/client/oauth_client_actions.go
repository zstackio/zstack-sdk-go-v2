// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateOAuthClient 更新OAuthClient
func (cli *ZSClient) UpdateOAuthClient(uuid string, params param.UpdateOAuthClientParam) (*view.UpdateOAuthClientEventView, error) {
	resp := view.UpdateOAuthClientEventView{}
	if err := cli.Put("v1/update/oauth2/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

