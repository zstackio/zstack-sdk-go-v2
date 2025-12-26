// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2VirtualIDGroup creates IAM2VirtualIDGroup
func (cli *ZSClient) CreateIAM2VirtualIDGroup(params param.CreateIAM2VirtualIDGroupParam) (*view.CreateIAM2VirtualIDGroupEventView, error) {
	resp := view.CreateIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
