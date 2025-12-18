// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddModel 操作AddModel
func (cli *ZSClient) AddModel(params param.AddModelParam) (*view.AddModelEventView, error) {
	resp := view.AddModelEventView{}
	if err := cli.Post("v1/ai/models", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

