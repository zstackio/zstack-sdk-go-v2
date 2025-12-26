// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIAM2VirtualIDsToProject adds IAM2VirtualIDsToProject
func (cli *ZSClient) AddIAM2VirtualIDsToProject(params param.AddIAM2VirtualIDsToProjectParam) (*view.AddIAM2VirtualIDsToProjectEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{projectUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
