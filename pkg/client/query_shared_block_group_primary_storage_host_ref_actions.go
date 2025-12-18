// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySharedBlockGroupPrimaryStorageHostRef queries SharedBlockGroupPrimaryStorageHostRef list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorageHostRef(params param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	return resp, cli.List("v1/sharedblock-group/host-refs", &params, &resp)
}
