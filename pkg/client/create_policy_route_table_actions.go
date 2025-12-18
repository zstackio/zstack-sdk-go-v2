// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePolicyRouteTable creates PolicyRouteTable
func (cli *ZSClient) CreatePolicyRouteTable(params param.CreatePolicyRouteTableParam) (*view.CreatePolicyRouteTableEventView, error) {
	resp := view.CreatePolicyRouteTableEventView{}
	if err := cli.Post("v1/policy-routes/tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
