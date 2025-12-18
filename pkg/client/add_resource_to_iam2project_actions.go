// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddResourceToIAM2Project 操作AddResourceToIAM2Project
func (cli *ZSClient) AddResourceToIAM2Project(params param.AddResourceToIAM2ProjectParam) (*view.AddResourceToIAM2ProjectEventView, error) {
	resp := view.AddResourceToIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects/add/resource/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

