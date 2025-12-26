// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddPreconfigurationTemplate adds PreconfigurationTemplate
func (cli *ZSClient) AddPreconfigurationTemplate(params param.AddPreconfigurationTemplateParam) (*view.AddPreconfigurationTemplateEventView, error) {
	resp := view.AddPreconfigurationTemplateEventView{}
	if err := cli.Post("v1/baremetal/preconfigurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
