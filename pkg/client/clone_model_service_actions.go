// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CloneModelService operates on ModelService
func (cli *ZSClient) CloneModelService(params param.CloneModelServiceParam) (*view.CloneModelServiceEventView, error) {
	resp := view.CloneModelServiceEventView{}
	if err := cli.Post("v1/ai/model-services/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
