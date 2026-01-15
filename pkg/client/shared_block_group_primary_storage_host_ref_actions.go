// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlockGroupPrimaryStorageHostRef queries SharedBlockGroupPrimaryStorageHostRef list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorageHostRef(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	return resp, cli.List("v1/sharedblock-group/host-refs", params, &resp)
}

// PageSharedBlockGroupPrimaryStorageHostRef Pagination
func (cli *ZSClient) PageSharedBlockGroupPrimaryStorageHostRef(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, int, error) {
	var sharedBlockGroupPrimaryStorageHostRefs []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	total, err := cli.Page("v1/sharedblock-group/host-refs", params, &sharedBlockGroupPrimaryStorageHostRefs)
	return sharedBlockGroupPrimaryStorageHostRefs, total, err
}
