// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcSharedQos queries VpcSharedQos list
func (cli *ZSClient) QueryVpcSharedQos(params *param.QueryParam) ([]view.VpcSharedQosInventoryView, error) {
	var resp []view.VpcSharedQosInventoryView
	return resp, cli.List("v1/vips/sharedqos", params, &resp)
}
