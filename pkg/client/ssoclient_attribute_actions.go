// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSSOClientAttribute updates SSOClientAttribute
func (cli *ZSClient) UpdateSSOClientAttribute(uuid string, params param.UpdateSSOClientAttributeParam) (*view.SSOClientAttributeInventoryView, error) {
	var resp view.UpdateSSOClientAttributeEventView
	if err := cli.Put("v1/sso/client/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
