// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcHaGroup creates VpcHaGroup
func (cli *ZSClient) CreateVpcHaGroup(params param.CreateVpcHaGroupParam) (*view.CreateVpcHaGroupEventView, error) {
	resp := view.CreateVpcHaGroupEventView{}
	if err := cli.Post("v1/vpc/hagroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
