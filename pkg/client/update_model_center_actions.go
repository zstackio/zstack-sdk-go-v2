// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateModelCenter updates ModelCenter
func (cli *ZSClient) UpdateModelCenter(uuid string, params param.UpdateModelCenterParam) (*view.UpdateModelCenterEventView, error) {
	resp := view.UpdateModelCenterEventView{}
	if err := cli.Put("v1/ai/model-centers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
