// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePolicyRouteTable creates PolicyRouteTable
func (cli *ZSClient) CreatePolicyRouteTable(params param.CreatePolicyRouteTableParam) (*view.CreatePolicyRouteTableEventView, error) {
	resp := view.CreatePolicyRouteTableEventView{}
	if err := cli.Post("v1/policy-routes/tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
