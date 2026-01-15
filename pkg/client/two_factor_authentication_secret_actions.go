// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetTwoFactorAuthenticationSecret gets TwoFactorAuthenticationSecret by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationSecret(uuid string) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	var resp view.GetTwoFactorAuthenticationSecretView
	if err := cli.Get("v1/twofactorauthentication/secret", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ResetTwoFactorAuthenticationSecret operates on TwoFactorAuthenticationSecret
func (cli *ZSClient) ResetTwoFactorAuthenticationSecret(uuid string, params param.ResetTwoFactorAuthenticationSecretParam) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	resp := view.TwoFactorAuthenticationSecretInventoryView{}
	if err := cli.Put("v1/twofactorauthentication/secrets", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
