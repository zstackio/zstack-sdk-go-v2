// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2VirtualIDsToOrganization 操作AddIAM2VirtualIDsToOrganization
func (cli *ZSClient) AddIAM2VirtualIDsToOrganization(params param.AddIAM2VirtualIDsToOrganizationParam) (*view.AddIAM2VirtualIDsToOrganizationEventView, error) {
	resp := view.AddIAM2VirtualIDsToOrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations/{organizationUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

