// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTwoFactorAuthenticationState gets TwoFactorAuthenticationState by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationState(uuid string) (*view.GetTwoFactorAuthenticationStateView, error) {
	var resp view.GetTwoFactorAuthenticationStateView
	if err := cli.Get("v1/twofactorauthentication/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
