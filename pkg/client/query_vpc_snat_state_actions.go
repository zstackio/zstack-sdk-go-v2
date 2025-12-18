// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcSnatState queries VpcSnatState list
func (cli *ZSClient) QueryVpcSnatState(params param.QueryParam) ([]view.VpcSnatStateInventoryView, error) {
	var resp []view.VpcSnatStateInventoryView
	return resp, cli.List("v1/vpc/virtual-routers/networkservicestate/snat", &params, &resp)
}
