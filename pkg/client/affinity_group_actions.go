// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
	return cli.Delete("v1/affinity-groups", uuid, string(deleteMode))
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

func (cli *ZSClient) GetAffinityGroup(uuid string) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.Get("v1/affinity-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
