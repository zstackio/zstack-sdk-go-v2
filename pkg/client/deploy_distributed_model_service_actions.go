// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeployDistributedModelService 操作DeployDistributedModelService
func (cli *ZSClient) DeployDistributedModelService(uuid string, params param.DeployDistributedModelServiceParam) (*view.DeployDistributedModelServiceEventView, error) {
	resp := view.DeployDistributedModelServiceEventView{}
	if err := cli.Put("v1/ai/model-services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

