// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEip queries Eip list
func (cli *ZSClient) QueryEip(params param.QueryParam) ([]view.EipInventoryView, error) {
	var resp []view.EipInventoryView
	return resp, cli.List("v1/eips", &params, &resp)
}
