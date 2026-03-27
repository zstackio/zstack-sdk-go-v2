// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAffinityGroup updates AffinityGroup
func (cli *ZSClient) UpdateAffinityGroup(ctx context.Context, uuid string, params param.UpdateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/affinity-groups", uuid, "", map[string]interface{}{
		"updateAffinityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAffinityGroup deletes AffinityGroup
func (cli *ZSClient) DeleteAffinityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/affinity-groups", uuid, string(deleteMode))
}
// CreateAffinityGroup creates AffinityGroup
func (cli *ZSClient) CreateAffinityGroup(ctx context.Context, params param.CreateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post(ctx, "v1/affinity-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAffinityGroup queries AffinityGroup list
func (cli *ZSClient) QueryAffinityGroup(ctx context.Context, params *param.QueryParam) ([]view.AffinityGroupInventoryView, error) {
	var resp []view.AffinityGroupInventoryView
	return resp, cli.List(ctx, "v1/affinity-groups", params, &resp)
}

func (cli *ZSClient) GetAffinityGroup(ctx context.Context, uuid string) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.Get(ctx, "v1/affinity-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAffinityGroup Pagination
func (cli *ZSClient) PageAffinityGroup(ctx context.Context, params *param.QueryParam) ([]view.AffinityGroupInventoryView, int, error) {
	var affinityGroups []view.AffinityGroupInventoryView
	total, err := cli.Page(ctx, "v1/affinity-groups", params, &affinityGroups)
	return affinityGroups, total, err
}
