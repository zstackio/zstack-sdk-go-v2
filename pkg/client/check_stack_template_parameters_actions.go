// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckStackTemplateParameters operates on CheckStackTemplateParameters
func (cli *ZSClient) CheckStackTemplateParameters(params param.CheckStackTemplateParametersParam) (*view.CheckStackTemplateParametersView, error) {
	resp := view.CheckStackTemplateParametersView{}
	if err := cli.Post("v1/cloudformation/stack/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
