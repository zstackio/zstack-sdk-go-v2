// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryImageGroupRef queries ImageGroupRef list
func (cli *ZSClient) QueryImageGroupRef(params param.QueryParam) ([]view.ImageGroupRefInventoryView, error) {
	var resp []view.ImageGroupRefInventoryView
	return resp, cli.List("v1/imagegrouprefs", &params, &resp)
}
