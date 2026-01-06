// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(params param.CreateCasClientParam) (*view.CasClientInventoryView, error) {
	var resp view.CreateCasClientEventView
	if err := cli.Post("v1/create/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(uuid string, params param.UpdateCasClientParam) (*view.CasClientInventoryView, error) {
	var resp view.UpdateCasClientEventView
	if err := cli.Put("v1/update/cas/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
