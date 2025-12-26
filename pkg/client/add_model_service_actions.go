// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddModelService adds ModelService
func (cli *ZSClient) AddModelService(params param.AddModelServiceParam) (*view.AddModelServiceEventView, error) {
	resp := view.AddModelServiceEventView{}
	if err := cli.Post("v1/ai/model-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
