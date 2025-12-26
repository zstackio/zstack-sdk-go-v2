// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2VirtualID creates IAM2VirtualID
func (cli *ZSClient) CreateIAM2VirtualID(params param.CreateIAM2VirtualIDParam) (*view.CreateIAM2VirtualIDEventView, error) {
	resp := view.CreateIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
