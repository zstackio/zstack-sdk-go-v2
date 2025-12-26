// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DeployModelService operates on DeployModelService
func (cli *ZSClient) DeployModelService(uuid string, params param.DeployModelServiceParam) (*view.DeployModelServiceEventView, error) {
	resp := view.DeployModelServiceEventView{}
	if err := cli.Put("v1/ai/model-services/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
