// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RefreshSSOServerToken operates on SSOServerToken
func (cli *ZSClient) RefreshSSOServerToken(uuid string, params param.RefreshSSOServerTokenParam) (*view.SSOServerTokenInventoryView, error) {
	var resp view.RefreshSSOServerTokenEventView
	if err := cli.Put("v1/sso/server/token/refresh", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
