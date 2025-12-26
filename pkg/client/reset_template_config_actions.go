// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResetTemplateConfig operates on ResetTemplateConfig
func (cli *ZSClient) ResetTemplateConfig(uuid string, params param.ResetTemplateConfigParam) (*view.ResetTemplateConfigEventView, error) {
	resp := view.ResetTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
