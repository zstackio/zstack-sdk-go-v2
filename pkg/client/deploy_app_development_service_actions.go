// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DeployAppDevelopmentService operates on DeployAppDevelopmentService
func (cli *ZSClient) DeployAppDevelopmentService(uuid string, params param.DeployAppDevelopmentServiceParam) (*view.DeployAppDevelopmentServiceEventView, error) {
	resp := view.DeployAppDevelopmentServiceEventView{}
	if err := cli.Put("v1/ai/model-services/app/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
