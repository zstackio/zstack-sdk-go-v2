// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2OrganizationAttribute updates IAM2OrganizationAttribute
func (cli *ZSClient) UpdateIAM2OrganizationAttribute(uuid string, params param.UpdateIAM2OrganizationAttributeParam) (*view.UpdateIAM2OrganizationAttributeEventView, error) {
	resp := view.UpdateIAM2OrganizationAttributeEventView{}
	if err := cli.Put("v1/iam2/organizations/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
