// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTwoFactorAuthentication queries TwoFactorAuthentication list
func (cli *ZSClient) QueryTwoFactorAuthentication(params *param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, error) {
	var resp []view.TwoFactorAuthenticationInventoryView
	return resp, cli.List("v1/twofactorauthentication/secrets", params, &resp)
}
