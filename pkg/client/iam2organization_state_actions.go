// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeIAM2OrganizationState 操作IAM2OrganizationState
func (cli *ZSClient) ChangeIAM2OrganizationState(uuid string, params param.ChangeIAM2OrganizationStateParam) (*view.ChangeIAM2OrganizationStateEventView, error) {
	resp := view.ChangeIAM2OrganizationStateEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

