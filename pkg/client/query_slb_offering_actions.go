// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySlbOffering queries SlbOffering list
func (cli *ZSClient) QuerySlbOffering(params *param.QueryParam) ([]view.SlbOfferingInventoryView, error) {
	var resp []view.SlbOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/slb", params, &resp)
}
