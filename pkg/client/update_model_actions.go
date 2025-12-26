// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(uuid string, params param.UpdateModelParam) (*view.UpdateModelEventView, error) {
	resp := view.UpdateModelEventView{}
	if err := cli.Put("v1/ai/models/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
