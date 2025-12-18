// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2VirtualIDGroupToProjects 操作AddIAM2VirtualIDGroupToProjects
func (cli *ZSClient) AddIAM2VirtualIDGroupToProjects(params param.AddIAM2VirtualIDGroupToProjectsParam) (*view.AddIAM2VirtualIDGroupToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDGroupToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

