// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeployModelService 操作DeployModelService
func (cli *ZSClient) DeployModelService(uuid string, params param.DeployModelServiceParam) (*view.DeployModelServiceEventView, error) {
	resp := view.DeployModelServiceEventView{}
	if err := cli.Put("v1/ai/model-services/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

