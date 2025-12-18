// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryZBox queries ZBox list
func (cli *ZSClient) QueryZBox(params param.QueryParam) ([]view.ZBoxInventoryView, error) {
	var resp []view.ZBoxInventoryView
	return resp, cli.List("v1/zbox", &params, &resp)
}
