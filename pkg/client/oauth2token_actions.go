// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetOAuth2Token 获取OAuth2Token详情
func (cli *ZSClient) GetOAuth2Token(uuid string) (*view.GetOAuth2TokenView, error) {
	var resp view.GetOAuth2TokenView
	if err := cli.Get("v1/get/oauth2/token", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

