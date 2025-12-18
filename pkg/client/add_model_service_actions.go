// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddModelService 操作AddModelService
func (cli *ZSClient) AddModelService(params param.AddModelServiceParam) (*view.AddModelServiceEventView, error) {
	resp := view.AddModelServiceEventView{}
	if err := cli.Post("v1/ai/model-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

