// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DecodeStackTemplate operates on DecodeStackTemplate
func (cli *ZSClient) DecodeStackTemplate(params param.DecodeStackTemplateParam) (*view.DecodeStackTemplateView, error) {
	resp := view.DecodeStackTemplateView{}
	if err := cli.Post("v1/cloudformation/stack/preview/resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
