// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTwoFactorAuthentication queries TwoFactorAuthentication list
func (cli *ZSClient) QueryTwoFactorAuthentication(params *param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, error) {
	var resp []view.TwoFactorAuthenticationInventoryView
	return resp, cli.List("v1/twofactorauthentication/secrets", params, &resp)
}

// PageTwoFactorAuthentication Pagination
func (cli *ZSClient) PageTwoFactorAuthentication(params *param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, int, error) {
	var twoFactorAuthentications []view.TwoFactorAuthenticationInventoryView
	total, err := cli.Page("v1/twofactorauthentication/secrets", params, &twoFactorAuthentications)
	return twoFactorAuthentications, total, err
}
