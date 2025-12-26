// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevertTemplateConfig operates on RevertTemplateConfig
func (cli *ZSClient) RevertTemplateConfig(uuid string, params param.RevertTemplateConfigParam) (*view.RevertTemplateConfigEventView, error) {
	resp := view.RevertTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
