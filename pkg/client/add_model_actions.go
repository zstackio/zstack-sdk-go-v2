// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddModel adds Model
func (cli *ZSClient) AddModel(params param.AddModelParam) (*view.AddModelEventView, error) {
	resp := view.AddModelEventView{}
	if err := cli.Post("v1/ai/models", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
