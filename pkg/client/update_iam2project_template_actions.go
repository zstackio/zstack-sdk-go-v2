// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2ProjectTemplate updates IAM2ProjectTemplate
func (cli *ZSClient) UpdateIAM2ProjectTemplate(uuid string, params param.UpdateIAM2ProjectTemplateParam) (*view.UpdateIAM2ProjectTemplateEventView, error) {
	resp := view.UpdateIAM2ProjectTemplateEventView{}
	if err := cli.Put("v1/iam2/projects/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
