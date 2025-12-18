// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePolicyRouteTableRouteEntry creates PolicyRouteTableRouteEntry
func (cli *ZSClient) CreatePolicyRouteTableRouteEntry(params param.CreatePolicyRouteTableRouteEntryParam) (*view.CreatePolicyRouteTableRouteEntryEventView, error) {
	resp := view.CreatePolicyRouteTableRouteEntryEventView{}
	if err := cli.Post("v1/policy-routes/routes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
