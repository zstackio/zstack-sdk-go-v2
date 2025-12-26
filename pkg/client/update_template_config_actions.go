// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateTemplateConfig updates TemplateConfig
func (cli *ZSClient) UpdateTemplateConfig(uuid string, params param.UpdateTemplateConfigParam) (*view.UpdateTemplateConfigEventView, error) {
	resp := view.UpdateTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
