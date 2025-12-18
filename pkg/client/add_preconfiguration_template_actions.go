// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddPreconfigurationTemplate 操作AddPreconfigurationTemplate
func (cli *ZSClient) AddPreconfigurationTemplate(params param.AddPreconfigurationTemplateParam) (*view.AddPreconfigurationTemplateEventView, error) {
	resp := view.AddPreconfigurationTemplateEventView{}
	if err := cli.Post("v1/baremetal/preconfigurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

