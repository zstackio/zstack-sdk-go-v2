// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryModelCenter queries ModelCenter list
func (cli *ZSClient) QueryModelCenter(params param.QueryParam) ([]view.ModelCenterInventoryView, error) {
	var resp []view.ModelCenterInventoryView
	return resp, cli.List("v1/ai/model-centers", &params, &resp)
}
