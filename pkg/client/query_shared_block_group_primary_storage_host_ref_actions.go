// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySharedBlockGroupPrimaryStorageHostRef queries SharedBlockGroupPrimaryStorageHostRef list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorageHostRef(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	return resp, cli.List("v1/sharedblock-group/host-refs", params, &resp)
}
