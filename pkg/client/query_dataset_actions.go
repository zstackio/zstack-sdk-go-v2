// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDataset queries Dataset list
func (cli *ZSClient) QueryDataset(params param.QueryParam) ([]view.DatasetInventoryView, error) {
	var resp []view.DatasetInventoryView
	return resp, cli.List("v1/ai/datasets", &params, &resp)
}
