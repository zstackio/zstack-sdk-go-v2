// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVRouterRouterId 操作SetVRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post("v1/routerArea/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

