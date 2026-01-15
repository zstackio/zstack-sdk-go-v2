// Copyright (c) ZStack.io, Inc.

package client

import (
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

// PageAccountResourceRef Pagination
func (cli *ZSClient) PageAccountResourceRef(params *param.QueryParam) ([]view.AccountResourceRefInventoryView, int, error) {
	var accountResourceRefs []view.AccountResourceRefInventoryView
	total, err := cli.Page("v1/accounts/resources/refs", params, &accountResourceRefs)
	return accountResourceRefs, total, err
}
