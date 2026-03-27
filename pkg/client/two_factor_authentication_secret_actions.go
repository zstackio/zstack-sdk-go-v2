// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetTwoFactorAuthenticationSecret gets TwoFactorAuthenticationSecret by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationSecret(ctx context.Context) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	var resp view.GetTwoFactorAuthenticationSecretView
	if err := cli.GetWithRespKey(ctx, "v1/twofactorauthentication/secret", "", "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ResetTwoFactorAuthenticationSecret operates on TwoFactorAuthenticationSecret
func (cli *ZSClient) ResetTwoFactorAuthenticationSecret(ctx context.Context, params param.ResetTwoFactorAuthenticationSecretParam) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	resp := view.TwoFactorAuthenticationSecretInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/twofactorauthentication/secrets", "", "", map[string]interface{}{
		"resetTwoFactorAuthenticationSecret": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
