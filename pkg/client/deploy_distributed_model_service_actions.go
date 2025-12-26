// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DeployDistributedModelService operates on DeployDistributedModelService
func (cli *ZSClient) DeployDistributedModelService(uuid string, params param.DeployDistributedModelServiceParam) (*view.DeployDistributedModelServiceEventView, error) {
	resp := view.DeployDistributedModelServiceEventView{}
	if err := cli.Put("v1/ai/model-services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
