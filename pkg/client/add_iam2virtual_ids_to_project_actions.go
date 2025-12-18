// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2VirtualIDsToProject 操作AddIAM2VirtualIDsToProject
func (cli *ZSClient) AddIAM2VirtualIDsToProject(params param.AddIAM2VirtualIDsToProjectParam) (*view.AddIAM2VirtualIDsToProjectEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectEventView{}
	if err := cli.Post("v1/iam2/projects/{projectUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

