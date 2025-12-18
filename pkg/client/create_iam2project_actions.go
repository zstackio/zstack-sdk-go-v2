// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2Project creates IAM2Project
func (cli *ZSClient) CreateIAM2Project(params param.CreateIAM2ProjectParam) (*view.CreateIAM2ProjectEventView, error) {
	resp := view.CreateIAM2ProjectEventView{}
	if err := cli.Post("v1/iam2/projects", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
