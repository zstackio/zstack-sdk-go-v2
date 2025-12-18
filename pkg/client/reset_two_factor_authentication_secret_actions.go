// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResetTwoFactorAuthenticationSecret operates on ResetTwoFactorAuthenticationSecret
func (cli *ZSClient) ResetTwoFactorAuthenticationSecret(uuid string, params param.ResetTwoFactorAuthenticationSecretParam) (*view.ResetTwoFactorAuthenticationSecretEventView, error) {
	resp := view.ResetTwoFactorAuthenticationSecretEventView{}
	if err := cli.Put("v1/twofactorauthentication/secrets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
