// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCdpPolicy creates CdpPolicy
func (cli *ZSClient) CreateCdpPolicy(params param.CreateCdpPolicyParam) (*view.CdpPolicyInventoryView, error) {
	var resp view.CreateCdpPolicyEventView
	if err := cli.Post("v1/cdp-backup-storage/policy", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCdpPolicy deletes CdpPolicy
func (cli *ZSClient) DeleteCdpPolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-backup-storage/policy", uuid, string(deleteMode))
}
// QueryCdpPolicy queries CdpPolicy list
func (cli *ZSClient) QueryCdpPolicy(params *param.QueryParam) ([]view.CdpPolicyInventoryView, error) {
	var resp []view.CdpPolicyInventoryView
	return resp, cli.List("v1/cdp-backup-storage/policy", params, &resp)
}

func (cli *ZSClient) GetCdpPolicy(uuid string) (*view.CdpPolicyInventoryView, error) {
	var resp view.CdpPolicyInventoryView
	if err := cli.Get("v1/cdp-backup-storage/policy", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateCdpPolicy updates CdpPolicy
func (cli *ZSClient) UpdateCdpPolicy(uuid string, params param.UpdateCdpPolicyParam) (*view.CdpPolicyInventoryView, error) {
	var resp view.UpdateCdpPolicyEventView
	if err := cli.Put("v1/cdp-backup-storage/policy", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
