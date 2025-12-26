// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2ProjectTemplate creates IAM2ProjectTemplate
func (cli *ZSClient) CreateIAM2ProjectTemplate(params param.CreateIAM2ProjectTemplateParam) (*view.CreateIAM2ProjectTemplateEventView, error) {
	resp := view.CreateIAM2ProjectTemplateEventView{}
	if err := cli.Post("v1/iam2/projects/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
