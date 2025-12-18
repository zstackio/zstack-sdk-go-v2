// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(params param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List("v1/primary-storage/local-storage/resource-refs", &params, &resp)
}
