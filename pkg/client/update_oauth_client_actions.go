// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateOAuthClient updates OAuthClient
func (cli *ZSClient) UpdateOAuthClient(uuid string, params param.UpdateOAuthClientParam) (*view.UpdateOAuthClientEventView, error) {
	resp := view.UpdateOAuthClientEventView{}
	if err := cli.Put("v1/update/oauth2/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
