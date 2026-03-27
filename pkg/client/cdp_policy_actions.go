// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCdpPolicy creates CdpPolicy
func (cli *ZSClient) CreateCdpPolicy(ctx context.Context, params param.CreateCdpPolicyParam) (*view.CdpPolicyInventoryView, error) {
	resp := view.CdpPolicyInventoryView{}
	if err := cli.Post(ctx, "v1/cdp-backup-storage/policy", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteCdpPolicy deletes CdpPolicy
func (cli *ZSClient) DeleteCdpPolicy(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cdp-backup-storage/policy", uuid, string(deleteMode))
}
// QueryCdpPolicy queries CdpPolicy list
func (cli *ZSClient) QueryCdpPolicy(ctx context.Context, params *param.QueryParam) ([]view.CdpPolicyInventoryView, error) {
	var resp []view.CdpPolicyInventoryView
	return resp, cli.List(ctx, "v1/cdp-backup-storage/policy", params, &resp)
}

func (cli *ZSClient) GetCdpPolicy(ctx context.Context, uuid string) (*view.CdpPolicyInventoryView, error) {
	var resp view.CdpPolicyInventoryView
	if err := cli.Get(ctx, "v1/cdp-backup-storage/policy", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCdpPolicy Pagination
func (cli *ZSClient) PageCdpPolicy(ctx context.Context, params *param.QueryParam) ([]view.CdpPolicyInventoryView, int, error) {
	var cdpPolicies []view.CdpPolicyInventoryView
	total, err := cli.Page(ctx, "v1/cdp-backup-storage/policy", params, &cdpPolicies)
	return cdpPolicies, total, err
}
// UpdateCdpPolicy updates CdpPolicy
func (cli *ZSClient) UpdateCdpPolicy(ctx context.Context, uuid string, params param.UpdateCdpPolicyParam) (*view.CdpPolicyInventoryView, error) {
	resp := view.CdpPolicyInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-backup-storage/policy", uuid, "", map[string]interface{}{
		"updateCdpPolicy": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
