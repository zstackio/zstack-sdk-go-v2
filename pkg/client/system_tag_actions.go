// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSystemTag updates SystemTag
func (cli *ZSClient) UpdateSystemTag(uuid string, params param.UpdateSystemTagParam) (*view.SystemTagInventoryView, error) {
	var resp view.UpdateSystemTagEventView
	if err := cli.Put("v1/system-tags/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSystemTag creates SystemTag
func (cli *ZSClient) CreateSystemTag(params param.CreateSystemTagParam) (*view.SystemTagInventoryView, error) {
	var resp view.CreateSystemTagEventView
	if err := cli.Post("v1/system-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySystemTag queries SystemTag list
func (cli *ZSClient) QuerySystemTag(params *param.QueryParam) ([]view.SystemTagInventoryView, error) {
	var resp []view.SystemTagInventoryView
	return resp, cli.List("v1/system-tags", params, &resp)
}
