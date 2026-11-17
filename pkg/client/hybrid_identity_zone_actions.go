// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SyncIdentityFromRemote operates on IdentityFromRemote
func (cli *ZSClient) SyncIdentityFromRemote(uuid string) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.GetWithRespKey("v1/hybrid/identity-zone", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIdentityZoneFromLocal queries IdentityZoneFromLocal list
func (cli *ZSClient) QueryIdentityZoneFromLocal(params *param.QueryParam) ([]view.IdentityZoneInventoryView, error) {
	var resp []view.IdentityZoneInventoryView
	return resp, cli.List("v1/hybrid/identity-zone", params, &resp)
}

func (cli *ZSClient) GetIdentityZoneFromLocal(uuid string) (*view.IdentityZoneInventoryView, error) {
	var resp view.IdentityZoneInventoryView
	if err := cli.Get("v1/hybrid/identity-zone", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIdentityZoneFromLocal Pagination
func (cli *ZSClient) PageIdentityZoneFromLocal(params *param.QueryParam) ([]view.IdentityZoneInventoryView, int, error) {
	var identityZoneFromLocals []view.IdentityZoneInventoryView
	total, err := cli.Page("v1/hybrid/identity-zone", params, &identityZoneFromLocals)
	return identityZoneFromLocals, total, err
}

// GetIdentityZoneFromRemote gets IdentityZoneFromRemote by uuid
func (cli *ZSClient) GetIdentityZoneFromRemote() (*view.IdentityZonePropertyView, error) {
	var resp view.IdentityZonePropertyView
	if err := cli.GetWithRespKey("v1/hybrid/identity-zone/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIdentityZoneFromRemote adds IdentityZoneFromRemote
func (cli *ZSClient) AddIdentityZoneFromRemote(params param.AddIdentityZoneFromRemoteParam) (*view.IdentityZoneInventoryView, error) {
	resp := view.IdentityZoneInventoryView{}
	if err := cli.Post("v1/hybrid/identity-zone", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIdentityZoneInLocal deletes IdentityZoneInLocal
func (cli *ZSClient) DeleteIdentityZoneInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/identity-zone", uuid, string(deleteMode))
}
