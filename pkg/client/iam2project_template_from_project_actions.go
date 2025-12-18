// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2ProjectTemplateFromProject 创建IAM2ProjectTemplateFromProject
func (cli *ZSClient) CreateIAM2ProjectTemplateFromProject(params param.CreateIAM2ProjectTemplateFromProjectParam) (*view.CreateIAM2ProjectTemplateFromProjectEventView, error) {
	resp := view.CreateIAM2ProjectTemplateFromProjectEventView{}
	if err := cli.Post("v1/iam2/projects/templates/from/projects/{projectUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

