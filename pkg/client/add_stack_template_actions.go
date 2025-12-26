// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddStackTemplate adds StackTemplate
func (cli *ZSClient) AddStackTemplate(params param.AddStackTemplateParam) (*view.AddStackTemplateEventView, error) {
	resp := view.AddStackTemplateEventView{}
	if err := cli.Post("v1/cloudformation/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
