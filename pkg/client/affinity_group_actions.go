// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAffinityGroup updates AffinityGroup
func (cli *ZSClient) UpdateAffinityGroup(uuid string, params param.UpdateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	var resp view.UpdateAffinityGroupEventView
	if err := cli.Put("v1/affinity-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAffinityGroup deletes AffinityGroup
func (cli *ZSClient) DeleteAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups/{uuid}", uuid, string(deleteMode))
}
// CreateAffinityGroup creates AffinityGroup
func (cli *ZSClient) CreateAffinityGroup(params param.CreateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	var resp view.CreateAffinityGroupEventView
	if err := cli.Post("v1/affinity-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAffinityGroup queries AffinityGroup list
func (cli *ZSClient) QueryAffinityGroup(params *param.QueryParam) ([]view.AffinityGroupInventoryView, error) {
	var resp []view.AffinityGroupInventoryView
	return resp, cli.List("v1/affinity-groups", params, &resp)
}
