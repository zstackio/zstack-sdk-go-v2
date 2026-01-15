// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List("v1/primary-storage/local-storage/resource-refs", params, &resp)
}

// PageLocalStorageResourceRef Pagination
func (cli *ZSClient) PageLocalStorageResourceRef(params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, int, error) {
	var localStorageResourceRefs []view.LocalStorageResourceRefInventoryView
	total, err := cli.Page("v1/primary-storage/local-storage/resource-refs", params, &localStorageResourceRefs)
	return localStorageResourceRefs, total, err
}
