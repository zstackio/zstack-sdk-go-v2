// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIAM2VirtualIDsToOrganization adds IAM2VirtualIDsToOrganization
func (cli *ZSClient) AddIAM2VirtualIDsToOrganization(params param.AddIAM2VirtualIDsToOrganizationParam) (*view.AddIAM2VirtualIDsToOrganizationEventView, error) {
	resp := view.AddIAM2VirtualIDsToOrganizationEventView{}
	if err := cli.Post("v1/iam2/organizations/{organizationUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
