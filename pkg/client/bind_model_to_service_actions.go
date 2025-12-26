// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BindModelToService operates on BindModelToService
func (cli *ZSClient) BindModelToService(params param.BindModelToServiceParam) (*view.BindModelToServiceEventView, error) {
	resp := view.BindModelToServiceEventView{}
	if err := cli.Post("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
