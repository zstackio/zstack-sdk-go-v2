// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeployModelEvalService operates on DeployModelEvalService
func (cli *ZSClient) DeployModelEvalService(uuid string, params param.DeployModelEvalServiceParam) (*view.DeployModelEvalServiceEventView, error) {
	resp := view.DeployModelEvalServiceEventView{}
	if err := cli.Put("v1/ai/model-services/eval/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
