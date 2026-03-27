// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(ctx context.Context, params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/local-storage/resource-refs", params, &resp)
}

func (cli *ZSClient) GetLocalStorageResourceRef(ctx context.Context, uuid string) (*view.LocalStorageResourceRefInventoryView, error) {
	var resp view.LocalStorageResourceRefInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/local-storage/resource-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalStorageResourceRef Pagination
func (cli *ZSClient) PageLocalStorageResourceRef(ctx context.Context, params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, int, error) {
	var localStorageResourceRefs []view.LocalStorageResourceRefInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/local-storage/resource-refs", params, &localStorageResourceRefs)
	return localStorageResourceRefs, total, err
}
