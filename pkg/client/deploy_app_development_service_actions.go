// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeployAppDevelopmentService 操作DeployAppDevelopmentService
func (cli *ZSClient) DeployAppDevelopmentService(uuid string, params param.DeployAppDevelopmentServiceParam) (*view.DeployAppDevelopmentServiceEventView, error) {
	resp := view.DeployAppDevelopmentServiceEventView{}
	if err := cli.Put("v1/ai/model-services/app/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

