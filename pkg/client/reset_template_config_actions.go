// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResetTemplateConfig 操作ResetTemplateConfig
func (cli *ZSClient) ResetTemplateConfig(uuid string, params param.ResetTemplateConfigParam) (*view.ResetTemplateConfigEventView, error) {
	resp := view.ResetTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

