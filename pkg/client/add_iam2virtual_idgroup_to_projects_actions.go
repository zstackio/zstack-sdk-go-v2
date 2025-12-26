// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIAM2VirtualIDGroupToProjects adds IAM2VirtualIDGroupToProjects
func (cli *ZSClient) AddIAM2VirtualIDGroupToProjects(params param.AddIAM2VirtualIDGroupToProjectsParam) (*view.AddIAM2VirtualIDGroupToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDGroupToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
