// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIAM2VirtualIDInGroup gets IAM2VirtualIDInGroup by uuid
func (cli *ZSClient) GetIAM2VirtualIDInGroup(uuid string) (*view.GetIAM2VirtualIDInGroupView, error) {
	var resp view.GetIAM2VirtualIDInGroupView
	if err := cli.Get("v1/iam2/IAM2VirtualIDGroup/IAM2VirtualID", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
