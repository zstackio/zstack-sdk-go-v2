// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2ProjectAttribute updates IAM2ProjectAttribute
func (cli *ZSClient) UpdateIAM2ProjectAttribute(uuid string, params param.UpdateIAM2ProjectAttributeParam) (*view.UpdateIAM2ProjectAttributeEventView, error) {
	resp := view.UpdateIAM2ProjectAttributeEventView{}
	if err := cli.Put("v1/iam2/projects/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
