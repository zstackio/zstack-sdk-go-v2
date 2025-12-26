// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIAM2Organization updates IAM2Organization
func (cli *ZSClient) UpdateIAM2Organization(uuid string, params param.UpdateIAM2OrganizationParam) (*view.UpdateIAM2OrganizationEventView, error) {
	resp := view.UpdateIAM2OrganizationEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
