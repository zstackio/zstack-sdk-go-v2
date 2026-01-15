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
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Put("v1/affinity-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAffinityGroup deletes AffinityGroup
func (cli *ZSClient) DeleteAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups", uuid, string(deleteMode))
}
// CreateAffinityGroup creates AffinityGroup
func (cli *ZSClient) CreateAffinityGroup(params param.CreateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post("v1/affinity-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAffinityGroup queries AffinityGroup list
func (cli *ZSClient) QueryAffinityGroup(params *param.QueryParam) ([]view.AffinityGroupInventoryView, error) {
	var resp []view.AffinityGroupInventoryView
	return resp, cli.List("v1/affinity-groups", params, &resp)
}

// PageAffinityGroup Pagination
func (cli *ZSClient) PageAffinityGroup(params *param.QueryParam) ([]view.AffinityGroupInventoryView, int, error) {
	var affinityGroups []view.AffinityGroupInventoryView
	total, err := cli.Page("v1/affinity-groups", params, &affinityGroups)
	return affinityGroups, total, err
}
