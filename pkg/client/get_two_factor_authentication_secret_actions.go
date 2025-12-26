// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTwoFactorAuthenticationSecret gets TwoFactorAuthenticationSecret by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationSecret(uuid string) (*view.GetTwoFactorAuthenticationSecretView, error) {
	var resp view.GetTwoFactorAuthenticationSecretView
	if err := cli.Get("v1/twofactorauthentication/secret", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
