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

func (cli *ZSClient) GetSharedBlockGroupPrimaryStorageHostRef(uuid string) (*view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	if err := cli.Get("v1/sharedblock-group/host-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
