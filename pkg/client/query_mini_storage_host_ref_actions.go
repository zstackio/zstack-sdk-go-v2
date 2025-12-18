// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMiniStorageHostRef queries MiniStorageHostRef list
func (cli *ZSClient) QueryMiniStorageHostRef(params param.QueryParam) ([]view.MiniStorageHostRefInventoryView, error) {
	var resp []view.MiniStorageHostRefInventoryView
	return resp, cli.List("v1/primary-storage/mini/host-refs", &params, &resp)
}
