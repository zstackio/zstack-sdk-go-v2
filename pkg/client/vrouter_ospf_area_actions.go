// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVRouterOspfArea 创建VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(params param.CreateVRouterOspfAreaParam) (*view.CreateVRouterOspfAreaEventView, error) {
	resp := view.CreateVRouterOspfAreaEventView{}
	if err := cli.Post("v1/routerArea", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

