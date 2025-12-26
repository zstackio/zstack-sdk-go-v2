// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2ProjectTemplateFromProject creates IAM2ProjectTemplateFromProject
func (cli *ZSClient) CreateIAM2ProjectTemplateFromProject(params param.CreateIAM2ProjectTemplateFromProjectParam) (*view.CreateIAM2ProjectTemplateFromProjectEventView, error) {
	resp := view.CreateIAM2ProjectTemplateFromProjectEventView{}
	if err := cli.Post("v1/iam2/projects/templates/from/projects/{projectUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
