// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIAM2OrganizationState changes IAM2OrganizationState
func (cli *ZSClient) ChangeIAM2OrganizationState(uuid string, params param.ChangeIAM2OrganizationStateParam) (*view.ChangeIAM2OrganizationStateEventView, error) {
	resp := view.ChangeIAM2OrganizationStateEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
