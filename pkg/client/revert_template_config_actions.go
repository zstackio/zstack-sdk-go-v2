// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RevertTemplateConfig operates on RevertTemplateConfig
func (cli *ZSClient) RevertTemplateConfig(uuid string, params param.RevertTemplateConfigParam) (*view.RevertTemplateConfigEventView, error) {
	resp := view.RevertTemplateConfigEventView{}
	if err := cli.Put("v1/template-configurations/{templateUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
