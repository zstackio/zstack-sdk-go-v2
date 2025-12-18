// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTwoFactorAuthentication 查询TwoFactorAuthentication列表
func (cli *ZSClient) QueryTwoFactorAuthentication(params param.QueryParam) ([]view.QueryTwoFactorAuthenticationView, error) {
	var resp []view.QueryTwoFactorAuthenticationView
	return resp, cli.List("v1/twofactorauthentication/secrets", &params, &resp)
}

