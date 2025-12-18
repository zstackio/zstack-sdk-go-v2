// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateStackTemplate 更新StackTemplate
func (cli *ZSClient) UpdateStackTemplate(uuid string, params param.UpdateStackTemplateParam) (*view.UpdateStackTemplateEventView, error) {
	resp := view.UpdateStackTemplateEventView{}
	if err := cli.Put("v1/cloudformation/template/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

