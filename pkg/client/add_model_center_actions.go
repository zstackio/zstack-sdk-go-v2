// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddModelCenter adds ModelCenter
func (cli *ZSClient) AddModelCenter(params param.AddModelCenterParam) (*view.AddModelCenterEventView, error) {
	resp := view.AddModelCenterEventView{}
	if err := cli.Post("v1/ai/model-centers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
