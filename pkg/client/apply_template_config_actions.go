// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ApplyTemplateConfig 操作ApplyTemplateConfig
func (cli *ZSClient) ApplyTemplateConfig(uuid string, params param.ApplyTemplateConfigParam) (*view.ApplyTemplateConfigEventView, error) {
	resp := view.ApplyTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

