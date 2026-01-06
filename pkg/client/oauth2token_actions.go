// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetOAuth2Token gets OAuth2Token by uuid
func (cli *ZSClient) GetOAuth2Token(uuid string) (*view.OAuth2TokenInventoryView, error) {
	var resp view.GetOAuth2TokenView
	if err := cli.Get("v1/get/oauth2/token", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
