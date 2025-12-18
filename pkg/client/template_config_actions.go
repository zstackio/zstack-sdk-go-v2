// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateTemplateConfig 更新TemplateConfig
func (cli *ZSClient) UpdateTemplateConfig(uuid string, params param.UpdateTemplateConfigParam) (*view.UpdateTemplateConfigEventView, error) {
	resp := view.UpdateTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

