// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateModelService updates ModelService
func (cli *ZSClient) UpdateModelService(uuid string, params param.UpdateModelServiceParam) (*view.UpdateModelServiceEventView, error) {
	resp := view.UpdateModelServiceEventView{}
	if err := cli.Put("v1/ai/model-services/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
