// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIAM2ProjectsOfVirtualID gets IAM2ProjectsOfVirtualID by uuid
func (cli *ZSClient) GetIAM2ProjectsOfVirtualID(uuid string) (*view.GetIAM2ProjectsOfVirtualIDView, error) {
	var resp view.GetIAM2ProjectsOfVirtualIDView
	if err := cli.Get("v1/iam2/virtual-ids/projects", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
