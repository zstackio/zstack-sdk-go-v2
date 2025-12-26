// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePreconfigurationTemplate updates PreconfigurationTemplate
func (cli *ZSClient) UpdatePreconfigurationTemplate(uuid string, params param.UpdatePreconfigurationTemplateParam) (*view.UpdatePreconfigurationTemplateEventView, error) {
	resp := view.UpdatePreconfigurationTemplateEventView{}
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
