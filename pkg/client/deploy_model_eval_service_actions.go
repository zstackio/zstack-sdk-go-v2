// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DeployModelEvalService operates on DeployModelEvalService
func (cli *ZSClient) DeployModelEvalService(uuid string, params param.DeployModelEvalServiceParam) (*view.DeployModelEvalServiceEventView, error) {
	resp := view.DeployModelEvalServiceEventView{}
	if err := cli.Put("v1/ai/model-services/eval/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
