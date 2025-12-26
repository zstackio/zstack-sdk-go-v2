// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResetTwoFactorAuthenticationSecret operates on ResetTwoFactorAuthenticationSecret
func (cli *ZSClient) ResetTwoFactorAuthenticationSecret(uuid string, params param.ResetTwoFactorAuthenticationSecretParam) (*view.ResetTwoFactorAuthenticationSecretEventView, error) {
	resp := view.ResetTwoFactorAuthenticationSecretEventView{}
	if err := cli.Put("v1/twofactorauthentication/secrets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
