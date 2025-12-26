// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2ProjectFromTemplate creates IAM2ProjectFromTemplate
func (cli *ZSClient) CreateIAM2ProjectFromTemplate(params param.CreateIAM2ProjectFromTemplateParam) (*view.CreateIAM2ProjectFromTemplateEventView, error) {
	resp := view.CreateIAM2ProjectFromTemplateEventView{}
	if err := cli.Post("v1/iam2/projects/from/templates/{templateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
