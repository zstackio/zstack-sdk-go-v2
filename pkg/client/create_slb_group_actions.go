// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSlbGroup creates SlbGroup
func (cli *ZSClient) CreateSlbGroup(params param.CreateSlbGroupParam) (*view.CreateSlbGroupEventView, error) {
	resp := view.CreateSlbGroupEventView{}
	if err := cli.Post("v1/load-balancers/slb/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
