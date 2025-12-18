// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTwoFactorAuthenticationSecret 获取TwoFactorAuthenticationSecret详情
func (cli *ZSClient) GetTwoFactorAuthenticationSecret(uuid string) (*view.GetTwoFactorAuthenticationSecretView, error) {
	var resp view.GetTwoFactorAuthenticationSecretView
	if err := cli.Get("v1/twofactorauthentication/secret", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

