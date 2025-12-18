// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddStackTemplate 操作AddStackTemplate
func (cli *ZSClient) AddStackTemplate(params param.AddStackTemplateParam) (*view.AddStackTemplateEventView, error) {
	resp := view.AddStackTemplateEventView{}
	if err := cli.Post("v1/cloudformation/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

