// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2ProjectFromTemplate 创建IAM2ProjectFromTemplate
func (cli *ZSClient) CreateIAM2ProjectFromTemplate(params param.CreateIAM2ProjectFromTemplateParam) (*view.CreateIAM2ProjectFromTemplateEventView, error) {
	resp := view.CreateIAM2ProjectFromTemplateEventView{}
	if err := cli.Post("v1/iam2/projects/from/templates/{templateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

