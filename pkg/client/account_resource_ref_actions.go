// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAccountResourceRef queries AccountResourceRef list
func (cli *ZSClient) QueryAccountResourceRef(params *param.QueryParam) ([]view.AccountResourceRefInventoryView, error) {
	var resp []view.AccountResourceRefInventoryView
	return resp, cli.List("v1/accounts/resources/refs", params, &resp)
}

func (cli *ZSClient) GetAccountResourceRef(uuid string) (*view.AccountResourceRefInventoryView, error) {
	var resp view.AccountResourceRefInventoryView
	if err := cli.Get("v1/accounts/resources/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountResourceRef Pagination
func (cli *ZSClient) PageAccountResourceRef(params *param.QueryParam) ([]view.AccountResourceRefInventoryView, int, error) {
	var accountResourceRefs []view.AccountResourceRefInventoryView
	total, err := cli.Page("v1/accounts/resources/refs", params, &accountResourceRefs)
	return accountResourceRefs, total, err
}

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(accountUuid string, params param.ChangeResourceOwnerParam) (*view.AccountResourceRefInventoryView, error) {
	resp := view.AccountResourceRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/account/%s/resources", accountUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShareResource operates on ShareResource
func (cli *ZSClient) ShareResource(params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.PutWithRespKey("v1/accounts/resources/actions", "", "", map[string]interface{}{
		"shareResource": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeResourceSharing operates on RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.PutWithRespKey("v1/accounts/resources/actions", "", "", map[string]interface{}{
		"revokeResourceSharing": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
