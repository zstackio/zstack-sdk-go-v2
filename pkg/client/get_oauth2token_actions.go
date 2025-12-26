// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetOAuth2Token gets OAuth2Token by uuid
func (cli *ZSClient) GetOAuth2Token(uuid string) (*view.GetOAuth2TokenView, error) {
	var resp view.GetOAuth2TokenView
	if err := cli.Get("v1/get/oauth2/token", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
