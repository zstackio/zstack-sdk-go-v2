// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDataCenterFromLocal queries DataCenterFromLocal list
func (cli *ZSClient) QueryDataCenterFromLocal(params param.QueryParam) ([]view.DataCenterInventoryView, error) {
	var resp []view.DataCenterInventoryView
	return resp, cli.List("v1/hybrid/data-center", &params, &resp)
}
