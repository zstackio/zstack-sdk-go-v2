// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteTable queries PolicyRouteTable list
func (cli *ZSClient) QueryPolicyRouteTable(params param.QueryParam) ([]view.PolicyRouteTableInventoryView, error) {
	var resp []view.PolicyRouteTableInventoryView
	return resp, cli.List("v1/policy-routes/tables", &params, &resp)
}
