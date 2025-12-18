// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2Organization creates IAM2Organization
func (cli *ZSClient) CreateIAM2Organization(params param.CreateIAM2OrganizationParam) (*view.CreateIAM2OrganizationEventView, error) {
	resp := view.CreateIAM2OrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
