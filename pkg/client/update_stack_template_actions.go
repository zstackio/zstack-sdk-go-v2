// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateStackTemplate updates StackTemplate
func (cli *ZSClient) UpdateStackTemplate(uuid string, params param.UpdateStackTemplateParam) (*view.UpdateStackTemplateEventView, error) {
	resp := view.UpdateStackTemplateEventView{}
	if err := cli.Put("v1/cloudformation/template/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
