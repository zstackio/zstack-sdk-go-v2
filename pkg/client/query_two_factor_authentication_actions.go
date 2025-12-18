// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTwoFactorAuthentication queries TwoFactorAuthentication list
func (cli *ZSClient) QueryTwoFactorAuthentication(params param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, error) {
	var resp []view.TwoFactorAuthenticationInventoryView
	return resp, cli.List("v1/twofactorauthentication/secrets", &params, &resp)
}
