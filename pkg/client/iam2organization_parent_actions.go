// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeIAM2OrganizationParent 操作IAM2OrganizationParent
func (cli *ZSClient) ChangeIAM2OrganizationParent(uuid string, params param.ChangeIAM2OrganizationParentParam) (*view.ChangeIAM2OrganizationParentEventView, error) {
	resp := view.ChangeIAM2OrganizationParentEventView{}
	if err := cli.Put("v1/iam2/organizations/{parentUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

