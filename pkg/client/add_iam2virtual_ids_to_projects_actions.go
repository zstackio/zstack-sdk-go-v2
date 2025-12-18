// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2VirtualIDsToProjects 操作AddIAM2VirtualIDsToProjects
func (cli *ZSClient) AddIAM2VirtualIDsToProjects(params param.AddIAM2VirtualIDsToProjectsParam) (*view.AddIAM2VirtualIDsToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

