// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcSharedQos queries VpcSharedQos list
func (cli *ZSClient) QueryVpcSharedQos(params param.QueryParam) ([]view.VpcSharedQosInventoryView, error) {
	var resp []view.VpcSharedQosInventoryView
	return resp, cli.List("v1/vips/sharedqos", &params, &resp)
}
