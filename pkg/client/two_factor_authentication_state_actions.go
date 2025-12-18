// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTwoFactorAuthenticationState 获取TwoFactorAuthenticationState详情
func (cli *ZSClient) GetTwoFactorAuthenticationState(uuid string) (*view.GetTwoFactorAuthenticationStateView, error) {
	var resp view.GetTwoFactorAuthenticationStateView
	if err := cli.Get("v1/twofactorauthentication/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

