// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTwoFactorAuthentication queries TwoFactorAuthentication list
func (cli *ZSClient) QueryTwoFactorAuthentication(ctx context.Context, params *param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, error) {
	var resp []view.TwoFactorAuthenticationInventoryView
	return resp, cli.List(ctx, "v1/twofactorauthentication/secrets", params, &resp)
}

func (cli *ZSClient) GetTwoFactorAuthentication(ctx context.Context, uuid string) (*view.TwoFactorAuthenticationInventoryView, error) {
	var resp view.TwoFactorAuthenticationInventoryView
	if err := cli.Get(ctx, "v1/twofactorauthentication/secrets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTwoFactorAuthentication Pagination
func (cli *ZSClient) PageTwoFactorAuthentication(ctx context.Context, params *param.QueryParam) ([]view.TwoFactorAuthenticationInventoryView, int, error) {
	var twoFactorAuthentications []view.TwoFactorAuthenticationInventoryView
	total, err := cli.Page(ctx, "v1/twofactorauthentication/secrets", params, &twoFactorAuthentications)
	return twoFactorAuthentications, total, err
}
