// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVRouterOspfArea creates VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(params param.CreateVRouterOspfAreaParam) (*view.CreateVRouterOspfAreaEventView, error) {
	resp := view.CreateVRouterOspfAreaEventView{}
	if err := cli.Post("v1/routerArea", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
