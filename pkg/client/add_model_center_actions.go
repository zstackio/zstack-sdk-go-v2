// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddModelCenter 操作AddModelCenter
func (cli *ZSClient) AddModelCenter(params param.AddModelCenterParam) (*view.AddModelCenterEventView, error) {
	resp := view.AddModelCenterEventView{}
	if err := cli.Post("v1/ai/model-centers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

