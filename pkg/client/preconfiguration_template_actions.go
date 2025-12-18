// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePreconfigurationTemplate 更新PreconfigurationTemplate
func (cli *ZSClient) UpdatePreconfigurationTemplate(uuid string, params param.UpdatePreconfigurationTemplateParam) (*view.UpdatePreconfigurationTemplateEventView, error) {
	resp := view.UpdatePreconfigurationTemplateEventView{}
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

