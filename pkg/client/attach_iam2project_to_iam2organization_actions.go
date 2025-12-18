// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachIAM2ProjectToIAM2Organization operates on IAM2ProjectToIAM2Organization
func (cli *ZSClient) AttachIAM2ProjectToIAM2Organization(params param.AttachIAM2ProjectToIAM2OrganizationParam) (*view.AttachIAM2ProjectToIAM2OrganizationEventView, error) {
	resp := view.AttachIAM2ProjectToIAM2OrganizationEventView{}
	if err := cli.Post("v1/iam2/projects/{projectUuid}/iam2/organizations/{organizationUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
