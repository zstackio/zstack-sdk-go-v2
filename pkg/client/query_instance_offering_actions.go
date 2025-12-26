// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryInstanceOffering queries InstanceOffering list
func (cli *ZSClient) QueryInstanceOffering(params *param.QueryParam) ([]view.InstanceOfferingInventoryView, error) {
	var resp []view.InstanceOfferingInventoryView
	return resp, cli.List("v1/instance-offerings", params, &resp)
}
